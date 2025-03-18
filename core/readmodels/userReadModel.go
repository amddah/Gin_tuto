package readmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserReadModel represents the user document stored in MongoDB
type UserReadModel struct {
	ID        string             `bson:"_id"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Version   int                `bson:"version"`
	MongoID   primitive.ObjectID `bson:"mongo_id,omitempty"`
}
