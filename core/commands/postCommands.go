package commands

import (
	"time"

	"github.com/google/uuid"
)

// PostCommandTypes
const (
	CreatePostCommand CommandType = "CREATE_POST"
	UpdatePostCommand CommandType = "UPDATE_POST"
	DeletePostCommand CommandType = "DELETE_POST"
)

// CreatePostCmd represents a command to create a post
type CreatePostCmd struct {
	BaseCommand
	Title   string `json:"title"`
	Content string `json:"content"`
}

// NewCreatePostCommand creates a new command to create a post
func NewCreatePostCommand(title, content string) *CreatePostCmd {
	return &CreatePostCmd{
		BaseCommand: BaseCommand{
			ID:        uuid.New().String(),
			Type:      CreatePostCommand,
			Timestamp: time.Now(),
		},
		Title:   title,
		Content: content,
	}
}

// UpdatePostCmd represents a command to update a post
type UpdatePostCmd struct {
	BaseCommand
	PostID  string `json:"post_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// NewUpdatePostCommand creates a new command to update a post
func NewUpdatePostCommand(postID, title, content string) *UpdatePostCmd {
	return &UpdatePostCmd{
		BaseCommand: BaseCommand{
			ID:        uuid.New().String(),
			Type:      UpdatePostCommand,
			Timestamp: time.Now(),
		},
		PostID:  postID,
		Title:   title,
		Content: content,
	}
}

// DeletePostCmd represents a command to delete a post
type DeletePostCmd struct {
	BaseCommand
	PostID string `json:"post_id"`
}

// NewDeletePostCommand creates a new command to delete a post
func NewDeletePostCommand(postID string) *DeletePostCmd {
	return &DeletePostCmd{
		BaseCommand: BaseCommand{
			ID:        uuid.New().String(),
			Type:      DeletePostCommand,
			Timestamp: time.Now(),
		},
		PostID: postID,
	}
}
