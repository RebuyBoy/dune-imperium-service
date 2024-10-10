package repositories

import (
	"dune-imperium-service/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type GameRepository interface {
	GetAll(name string, name2 string) ([]models.Result, error)
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

func (g gameRepository) GetAll(string, string) ([]models.Result, error) {
	//TODO implement me
	panic("implement me")
}
