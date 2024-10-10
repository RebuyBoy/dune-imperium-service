package repositories

import (
	"dune-imperium-service/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type GameRepository interface {
	GetAll() ([]models.Result, error)
}

type gameRepository struct {
	collection *mongo.Collection
}

func NewGameRepository(db *mongo.Client) GameRepository {
	collection := db.Database("dune").Collection("results")
	return &gameRepository{
		collection: collection,
	}
}

func (g gameRepository) GetAll() ([]models.Result, error) {
	//TODO implement me
	panic("implement me")
}
