package events

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

// EventPublisher is responsible for publishing events
type EventPublisher interface {
	Publish(event Event) error
}

// RedisEventPublisher implements EventPublisher using Redis
type RedisEventPublisher struct {
	redisClient *redis.Client
	channelName string
}

// NewRedisEventPublisher creates a new Redis event publisher
func NewRedisEventPublisher(client *redis.Client, channel string) EventPublisher {
	return &RedisEventPublisher{
		redisClient: client,
		channelName: channel,
	}
}

// Publish publishes an event to Redis
func (p *RedisEventPublisher) Publish(event Event) error {
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return err
	}

	ctx := context.Background()
	err = p.redisClient.Publish(ctx, p.channelName, eventJSON).Err()
	if err != nil {
		return err
	}

	// Also store the event in a list for replay capability
	eventList := "events:" + string(event.GetType())
	err = p.redisClient.RPush(ctx, eventList, eventJSON).Err()
	if err != nil {
		log.Printf("Failed to store event in list: %v", err)
	}

	return nil
}
