package handlers

import (
	"encoding/json"
	repository "gin_api/adapter/Repository"
	"gin_api/core/events"
	"gin_api/core/readmodels"
	"log"
	"time"
)

// UserEventHandler handles user events
type UserEventHandler struct {
	userReadRepo *repository.MongoUserRepositoryImpl
}

// NewUserEventHandler creates a new user event handler
func NewUserEventHandler() *UserEventHandler {
	return &UserEventHandler{
		userReadRepo: repository.NewMongoUserRepository().(*repository.MongoUserRepositoryImpl),
	}
}

// HandleEvent handles an event
func (h *UserEventHandler) HandleEvent(event events.Event) error {
	switch event.GetType() {
	case events.UserCreatedEvent:
		return h.handleUserCreated(event)
	default:
		log.Printf("Unknown event type: %s", event.GetType())
		return nil
	}
}

// handleUserCreated handles a user created event
func (h *UserEventHandler) handleUserCreated(event events.Event) error {
	// Log the event
	payload := event.GetPayload()
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	log.Printf("User created event received: %s", string(jsonData))

	// Convert payload to map
	payloadMap, ok := payload.(map[string]interface{})
	if !ok {
		log.Printf("Failed to convert payload to map[string]interface{}")
		return nil
	}

	// Extract user information
	id, _ := payloadMap["id"].(string)
	name, _ := payloadMap["name"].(string)
	email, _ := payloadMap["email"].(string)

	// Create and save the user read model
	userReadModel := &readmodels.UserReadModel{
		ID:        id,
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}

	return h.userReadRepo.Save(userReadModel)
}
