package repository

import (
	"encoding/json"
	"gin_api/core/events"
	"gin_api/core/interfaces"
	"gin_api/core/models"
	"gin_api/initializer"
	"strings"

	"gorm.io/gorm"
)

// EventStoreRepositoryImpl implements the EventStoreRepository interface
type EventStoreRepositoryImpl struct {
	DB *gorm.DB
}

// NewEventStoreRepository creates a new event store repository
func NewEventStoreRepository() interfaces.EventStoreRepository {
	return &EventStoreRepositoryImpl{DB: initializer.DB}
}

// StoreEvent stores an event in the event store
func (repo *EventStoreRepositoryImpl) StoreEvent(event events.Event) error {
	// Convert event data to JSON
	eventData, err := json.Marshal(event.GetPayload())
	if err != nil {
		return err
	}

	// Determine aggregate type from event type
	// Convention: event type is like "USER_CREATED", "POST_UPDATED"
	// So we extract the first part before the underscore
	eventTypeStr := string(event.GetType())
	aggregateType := strings.Split(eventTypeStr, "_")[0]

	// Create event store entry
	eventStore := models.EventStore{
		EventID:       event.GetID(),
		EventType:     eventTypeStr,
		AggregateID:   event.GetAggregateID(),
		AggregateType: aggregateType,
		EventData:     string(eventData),
		Version:       1, // This should be incremented for each new event for the same aggregate
		CreatedAt:     event.GetTimestamp(),
	}

	// Store event
	return repo.DB.Create(&eventStore).Error
}

// GetEventsByAggregateID retrieves all events for a specific aggregate ID
func (repo *EventStoreRepositoryImpl) GetEventsByAggregateID(aggregateID string) ([]models.EventStore, error) {
	var events []models.EventStore
	err := repo.DB.Where("aggregate_id = ?", aggregateID).Order("version asc").Find(&events).Error
	return events, err
}

// GetEventsByType retrieves all events of a specific type
func (repo *EventStoreRepositoryImpl) GetEventsByType(eventType string) ([]models.EventStore, error) {
	var events []models.EventStore
	err := repo.DB.Where("event_type = ?", eventType).Order("created_at asc").Find(&events).Error
	return events, err
}

// GetAllEvents retrieves all events
func (repo *EventStoreRepositoryImpl) GetAllEvents() ([]models.EventStore, error) {
	var events []models.EventStore
	err := repo.DB.Order("created_at asc").Find(&events).Error
	return events, err
}
