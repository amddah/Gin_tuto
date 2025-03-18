package repository

import (
	"context"
	"gin_api/core/interfaces"
	"gin_api/core/readmodels"
	"gin_api/initializer"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoUserRepositoryImpl struct {
	collection *mongo.Collection
}

func NewMongoUserRepository() interfaces.UserReadRepository {
	collection := initializer.MongoDB.Collection("users")

	// Create indexes for faster lookup
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create index for email field
	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})

	if err != nil {
		// Log but don't fail - index might already exist
		// log.Printf("Error creating MongoDB index: %v", err)
	}

	return &MongoUserRepositoryImpl{collection: collection}
}

func (repo *MongoUserRepositoryImpl) FindByID(id string) (*readmodels.UserReadModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user readmodels.UserReadModel
	err := repo.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *MongoUserRepositoryImpl) FindAll() ([]readmodels.UserReadModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []readmodels.UserReadModel
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *MongoUserRepositoryImpl) Save(user *readmodels.UserReadModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}

	_, err := repo.collection.UpdateOne(ctx, filter, update, opts)
	return err
}

// FindByUUID is just an adapter for the existing code
func (repo *MongoUserRepositoryImpl) FindByUUID(uid uuid.UUID) (*readmodels.UserReadModel, error) {
	return repo.FindByID(uid.String())
}
