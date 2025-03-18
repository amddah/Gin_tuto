package handlers

import (
	repository "gin_api/adapter/Repository"
	"gin_api/core/commands"
	"gin_api/core/events"
	"gin_api/core/interfaces"

	"github.com/google/uuid"
)

// CommandHandler handles commands
type CommandHandler interface {
	Handle(command commands.Command) error
}

// UserCommandHandler handles user commands
type UserCommandHandler struct {
	publisher  events.EventPublisher
	eventStore interfaces.EventStoreRepository
}

// NewUserCommandHandler creates a new user command handler
func NewUserCommandHandler(publisher events.EventPublisher) *UserCommandHandler {
	return &UserCommandHandler{
		publisher:  publisher,
		eventStore: repository.NewEventStoreRepository(),
	}
}

// Handle handles a command
func (h *UserCommandHandler) Handle(command commands.Command) error {
	switch cmd := command.(type) {
	case *commands.CreateUserCmd:
		return h.handleCreateUser(cmd)
	default:
		return nil
	}
}

// handleCreateUser handles a create user command
func (h *UserCommandHandler) handleCreateUser(cmd *commands.CreateUserCmd) error {
	// Generate a new UUID for the user
	userID := uuid.New().String()

	// Create and publish event
	event := events.NewEvent(events.UserCreatedEvent, userID, map[string]interface{}{
		"id":    userID,
		"name":  cmd.Name,
		"email": cmd.Email,
	})

	// Store event in event store
	if err := h.eventStore.StoreEvent(event); err != nil {
		return err
	}

	// Publish event for subscribers
	return h.publisher.Publish(event)
}
