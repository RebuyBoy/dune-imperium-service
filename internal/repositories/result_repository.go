package repositories

import (
	"dune-imperium-service/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type ResultRepository interface {
	GetAll() ([]models.Result, error)
}

type resultRepository struct {
	collection *mongo.Collection
}

func NewResultRepository(db *mongo.Client) ResultRepository {
	collection := db.Database("dune").Collection("results")
	return &resultRepository{
		collection: collection,
	}
}

func (r resultRepository) GetAll() ([]models.Result, error) {
	//TODO implement me
	panic("implement me")
}
