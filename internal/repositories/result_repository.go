package repositories

import (
	"context"
	"dune-imperium-service/internal/models"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ResultRepository struct {
	collection *mongo.Collection
}

func NewResultRepository(db *mongo.Client) *ResultRepository {
	collection := db.Database("dune").Collection("results")
	return &ResultRepository{
		collection: collection,
	}
}

func (r ResultRepository) GetAll(ctx context.Context) ([]models.Result, error) {
	results := make([]models.Result, 0)

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return results, nil
		}
		return results, fmt.Errorf("failed to execute find query: %w", err)
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &results); err != nil {
		return results, fmt.Errorf("failed to decode results: %w", err)
	}

	return results, nil
}

func (r ResultRepository) Save(ctx context.Context, result models.Result) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, result)
	return err
}
