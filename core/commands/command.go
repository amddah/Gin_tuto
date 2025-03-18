package commands

import (
	"time"

	"github.com/google/uuid"
)

// CommandType represents the type of command
type CommandType string

// Command types
const (
	CreateUserCommand CommandType = "CREATE_USER"
	UpdateUserCommand CommandType = "UPDATE_USER"
	DeleteUserCommand CommandType = "DELETE_USER"
)

// Command is the base interface for all commands
type Command interface {
	GetID() string
	GetType() CommandType
	GetTimestamp() time.Time
}

// BaseCommand provides common functionality for all commands
type BaseCommand struct {
	ID        string      `json:"id"`
	Type      CommandType `json:"type"`
	Timestamp time.Time   `json:"timestamp"`
}

// GetID returns the command ID
func (c BaseCommand) GetID() string {
	return c.ID
}

// GetType returns the command type
func (c BaseCommand) GetType() CommandType {
	return c.Type
}

// GetTimestamp returns the timestamp of the command
func (c BaseCommand) GetTimestamp() time.Time {
	return c.Timestamp
}

// CreateUserCmd represents a command to create a user
type CreateUserCmd struct {
	BaseCommand
	Name  string `json:"name"`
	Email string `json:"email"`
}

// NewCreateUserCommand creates a new command to create a user
func NewCreateUserCommand(name, email string) *CreateUserCmd {
	return &CreateUserCmd{
		BaseCommand: BaseCommand{
			ID:        uuid.New().String(),
			Type:      CreateUserCommand,
			Timestamp: time.Now(),
		},
		Name:  name,
		Email: email,
	}
}
