package handlers

import (
	repository "gin_api/adapter/Repository"
	"gin_api/core/commands"
	"gin_api/core/events"
	"gin_api/core/interfaces"

	"github.com/google/uuid"
)

// PostCommandHandler handles post commands
type PostCommandHandler struct {
	publisher  events.EventPublisher
	eventStore interfaces.EventStoreRepository
}

// NewPostCommandHandler creates a new post command handler
func NewPostCommandHandler(publisher events.EventPublisher) *PostCommandHandler {
	return &PostCommandHandler{
		publisher:  publisher,
		eventStore: repository.NewEventStoreRepository(),
	}
}

// Handle handles a command
func (h *PostCommandHandler) Handle(command commands.Command) error {
	switch cmd := command.(type) {
	case *commands.CreatePostCmd:
		return h.handleCreatePost(cmd)
	case *commands.UpdatePostCmd:
		return h.handleUpdatePost(cmd)
	case *commands.DeletePostCmd:
		return h.handleDeletePost(cmd)
	default:
		return nil
	}
}

// handleCreatePost handles a create post command
func (h *PostCommandHandler) handleCreatePost(cmd *commands.CreatePostCmd) error {
	// Generate a new UUID for the post
	postID := uuid.New().String()

	// Create and publish event
	event := events.NewEvent(events.PostCreatedEvent, postID, map[string]interface{}{
		"id":      postID,
		"title":   cmd.Title,
		"content": cmd.Content,
	})

	// Store event in event store
	if err := h.eventStore.StoreEvent(event); err != nil {
		return err
	}

	// Publish event for subscribers
	return h.publisher.Publish(event)
}

// handleUpdatePost handles an update post command
func (h *PostCommandHandler) handleUpdatePost(cmd *commands.UpdatePostCmd) error {
	event := events.NewEvent(events.PostUpdatedEvent, cmd.PostID, map[string]interface{}{
		"id":      cmd.PostID,
		"title":   cmd.Title,
		"content": cmd.Content,
	})

	// Store event in event store
	if err := h.eventStore.StoreEvent(event); err != nil {
		return err
	}

	// Publish event for subscribers
	return h.publisher.Publish(event)
}

// handleDeletePost handles a delete post command
func (h *PostCommandHandler) handleDeletePost(cmd *commands.DeletePostCmd) error {
	event := events.NewEvent(events.PostDeletedEvent, cmd.PostID, map[string]interface{}{
		"id": cmd.PostID,
	})

	// Store event in event store
	if err := h.eventStore.StoreEvent(event); err != nil {
		return err
	}

	// Publish event for subscribers
	return h.publisher.Publish(event)
}
