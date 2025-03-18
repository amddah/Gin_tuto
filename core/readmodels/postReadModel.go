package readmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PostReadModel represents the post document stored in MongoDB
type PostReadModel struct {
	ID        string             `bson:"_id"` // Changed from uint to string and from "id" to "_id"
	Title     string             `bson:"title"`
	Content   string             `bson:"content"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Version   int                `bson:"version"`
	MongoID   primitive.ObjectID `bson:"mongo_id,omitempty"`
}
