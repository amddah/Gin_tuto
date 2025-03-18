package handlers

import (
	"encoding/json"
	repository "gin_api/adapter/Repository"
	"gin_api/core/events"
	"gin_api/core/readmodels"
	"log"
	"time"
)

// PostEventHandler handles post events
type PostEventHandler struct {
	postReadRepo *repository.MongoPostRepositoryImpl
}

// NewPostEventHandler creates a new post event handler
func NewPostEventHandler() *PostEventHandler {
	return &PostEventHandler{
		postReadRepo: repository.NewMongoPostRepository().(*repository.MongoPostRepositoryImpl),
	}
}

// HandleEvent handles an event
func (h *PostEventHandler) HandleEvent(event events.Event) error {
	switch event.GetType() {
	case events.PostCreatedEvent:
		return h.handlePostCreated(event)
	case events.PostUpdatedEvent:
		return h.handlePostUpdated(event)
	case events.PostDeletedEvent:
		return h.handlePostDeleted(event)
	default:
		log.Printf("Unknown event type: %s", event.GetType())
		return nil
	}
}

// handlePostCreated handles a post created event
func (h *PostEventHandler) handlePostCreated(event events.Event) error {
	payload := event.GetPayload()
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	log.Printf("Post created event received: %s", string(jsonData))

	// Convert payload to map
	payloadMap, ok := payload.(map[string]interface{})
	if !ok {
		log.Printf("Failed to convert payload to map[string]interface{}")
		return nil
	}

	// Extract post information
	id, _ := payloadMap["id"].(string)
	title, _ := payloadMap["title"].(string)
	content, _ := payloadMap["content"].(string)

	// Create and save the post read model
	postReadModel := &readmodels.PostReadModel{
		ID:        id,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}

	return h.postReadRepo.Save(postReadModel)
}

// handlePostUpdated handles a post updated event
func (h *PostEventHandler) handlePostUpdated(event events.Event) error {
	payload := event.GetPayload()
	payloadMap, ok := payload.(map[string]interface{})
	if !ok {
		return nil
	}

	id, _ := payloadMap["id"].(string)
	title, _ := payloadMap["title"].(string)
	content, _ := payloadMap["content"].(string)

	// Get existing post first
	existingPost, err := h.postReadRepo.FindByID(id)
	if err != nil {
		// If not found, create new
		postReadModel := &readmodels.PostReadModel{
			ID:        id,
			Title:     title,
			Content:   content,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Version:   1,
		}
		return h.postReadRepo.Save(postReadModel)
	}

	// Update existing post
	existingPost.Title = title
	existingPost.Content = content
	existingPost.UpdatedAt = time.Now()
	existingPost.Version++

	return h.postReadRepo.Save(existingPost)
}

// handlePostDeleted handles a post deleted event
func (h *PostEventHandler) handlePostDeleted(event events.Event) error {
	// In a real application, you might want to mark the post as deleted
	// or actually remove it from the read model
	// For simplicity, we'll just log the event
	payload := event.GetPayload()
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	log.Printf("Post deleted event received: %s", string(jsonData))
	return nil
}
