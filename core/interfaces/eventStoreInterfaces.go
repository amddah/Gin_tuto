package interfaces

import (
	"gin_api/core/events"
	"gin_api/core/models"
)

// EventStoreRepository defines the repository for storing events
type EventStoreRepository interface {
	StoreEvent(event events.Event) error
	GetEventsByAggregateID(aggregateID string) ([]models.EventStore, error)
	GetEventsByType(eventType string) ([]models.EventStore, error)
	GetAllEvents() ([]models.EventStore, error)
}
