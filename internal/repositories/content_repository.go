package repositories

import (
	"context"
	"dune-imperium-service/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ContentRepository struct {
	collection *mongo.Collection
}

func NewContentRepository(db *mongo.Client) *ContentRepository {
	collection := db.Database("dune").Collection("content")
	return &ContentRepository{
		collection: collection,
	}
}

func (r *ContentRepository) Save(ctx context.Context, content *models.GameContent) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := r.collection.InsertOne(ctx, content)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(string), nil
}

func (r *ContentRepository) FindById(ctx context.Context, id string) (*models.GameContent, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var content models.GameContent
	err := r.collection.FindOne(ctx, map[string]string{"_id": id}).Decode(&content)
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func (r *ContentRepository) FindByType(ctx context.Context, contentType models.ContentType) ([]*models.GameContent, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var content []*models.GameContent
	cursor, err := r.collection.Find(ctx, map[string]string{"type": string(contentType)})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &content); err != nil {
		return nil, err
	}

	return content, nil
}
