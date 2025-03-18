package models

import (
	"time"

	"gorm.io/gorm"
)

// EventStore represents an event in the event store
type EventStore struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	EventID       string         `gorm:"type:varchar(64);uniqueIndex" json:"event_id"`
	EventType     string         `gorm:"type:varchar(64);index" json:"event_type"`
	AggregateID   string         `gorm:"type:varchar(64);index" json:"aggregate_id"`
	AggregateType string         `gorm:"type:varchar(64);index" json:"aggregate_type"`
	EventData     string         `gorm:"type:text" json:"event_data"`
	Version       int            `gorm:"index" json:"version"`
	CreatedAt     time.Time      `json:"created_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
