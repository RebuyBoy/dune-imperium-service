package repositories

import "go.mongodb.org/mongo-driver/mongo"

type ContentRepository struct {
	collection *mongo.Collection
}

func NewContentRepository(db *mongo.Client) *ContentRepository {
	collection := db.Database("dune").Collection("content")
	return &ContentRepository{
		collection: collection,
	}
}
