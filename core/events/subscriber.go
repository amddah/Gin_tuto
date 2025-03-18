package events

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

// EventHandler handles incoming events
type EventHandler interface {
	HandleEvent(event Event) error
}

// EventSubscriber subscribes to events
type EventSubscriber interface {
	Subscribe(handler EventHandler)
	Close()
}

// RedisEventSubscriber implements EventSubscriber using Redis
type RedisEventSubscriber struct {
	redisClient *redis.Client
	channelName string
	pubsub      *redis.PubSub
	ctx         context.Context
	cancel      context.CancelFunc
}

// NewRedisEventSubscriber creates a new Redis event subscriber
func NewRedisEventSubscriber(client *redis.Client, channel string) EventSubscriber {
	ctx, cancel := context.WithCancel(context.Background())
	return &RedisEventSubscriber{
		redisClient: client,
		channelName: channel,
		ctx:         ctx,
		cancel:      cancel,
	}
}

// Subscribe subscribes to events and processes them with the given handler
func (s *RedisEventSubscriber) Subscribe(handler EventHandler) {
	s.pubsub = s.redisClient.Subscribe(s.ctx, s.channelName)

	// Start subscription in a goroutine
	go func() {
		channel := s.pubsub.Channel()
		for {
			select {
			case msg, ok := <-channel:
				if !ok {
					return
				}

				var baseEvent BaseEvent
				err := json.Unmarshal([]byte(msg.Payload), &baseEvent)
				if err != nil {
					log.Printf("Failed to deserialize event: %v", err)
					continue
				}

				err = handler.HandleEvent(baseEvent)
				if err != nil {
					log.Printf("Failed to handle event: %v", err)
				}

			case <-s.ctx.Done():
				return
			}
		}
	}()
}

// Close closes the subscription
func (s *RedisEventSubscriber) Close() {
	s.cancel()
	if s.pubsub != nil {
		s.pubsub.Close()
	}
}
