package events

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// EventType represents the type of event
type EventType string

// Command types
const (
	UserCreatedEvent EventType = "USER_CREATED"
	UserUpdatedEvent EventType = "USER_UPDATED"
	UserDeletedEvent EventType = "USER_DELETED"
)

// Event is the base interface for all events
type Event interface {
	GetID() string
	GetType() EventType
	GetAggregateID() string
	GetTimestamp() time.Time
	GetPayload() interface{}
}

// BaseEvent provides common functionality for all events
type BaseEvent struct {
	ID          string      `json:"id"`
	Type        EventType   `json:"type"`
	AggregateID string      `json:"aggregate_id"`
	Timestamp   time.Time   `json:"timestamp"`
	Payload     interface{} `json:"payload"`
}

// GetID returns the event ID
func (e BaseEvent) GetID() string {
	return e.ID
}

// GetType returns the event type
func (e BaseEvent) GetType() EventType {
	return e.Type
}

// GetAggregateID returns the aggregate ID
func (e BaseEvent) GetAggregateID() string {
	return e.AggregateID
}

// GetTimestamp returns the timestamp of the event
func (e BaseEvent) GetTimestamp() time.Time {
	return e.Timestamp
}

// GetPayload returns the event payload
func (e BaseEvent) GetPayload() interface{} {
	return e.Payload
}

// NewEvent creates a new event
func NewEvent(eventType EventType, aggregateID string, payload interface{}) Event {
	return BaseEvent{
		ID:          uuid.New().String(),
		Type:        eventType,
		AggregateID: aggregateID,
		Timestamp:   time.Now(),
		Payload:     payload,
	}
}

// Serialize converts an event to JSON
func Serialize(event Event) ([]byte, error) {
	return json.Marshal(event)
}

// Deserialize converts JSON to an event
func Deserialize(data []byte) (BaseEvent, error) {
	var event BaseEvent
	err := json.Unmarshal(data, &event)
	return event, err
}
