package repository

import (
	"context"
	"gin_api/core/interfaces"
	"gin_api/core/readmodels"
	"gin_api/initializer"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoPostRepositoryImpl implements the post read repository
type MongoPostRepositoryImpl struct {
	collection *mongo.Collection
}

// NewMongoPostRepository creates a new MongoDB post repository
func NewMongoPostRepository() interfaces.PostReadRepository {
	collection := initializer.MongoDB.Collection("posts")

	// Create indexes for faster lookup
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create index for title field for faster search
	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "title", Value: 1}},
		Options: options.Index().SetBackground(true),
	})

	if err != nil {
		// Log but don't fail - index might already exist
	}

	return &MongoPostRepositoryImpl{collection: collection}
}

// FindByID retrieves a post by ID
func (repo *MongoPostRepositoryImpl) FindByID(id string) (*readmodels.PostReadModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var post readmodels.PostReadModel
	err := repo.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// FindAll retrieves all posts
func (repo *MongoPostRepositoryImpl) FindAll() ([]readmodels.PostReadModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var posts []readmodels.PostReadModel
	if err := cursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

// Save stores a post in MongoDB
func (repo *MongoPostRepositoryImpl) Save(post *readmodels.PostReadModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": post.ID}
	update := bson.M{"$set": post}

	_, err := repo.collection.UpdateOne(ctx, filter, update, opts)
	return err
}
