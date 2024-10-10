package repositories

import (
	"dune-imperium-service/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetAll() ([]models.Result, error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Client) UserRepository {
	collection := db.Database("dune").Collection("results")
	return &userRepository{
		collection: collection,
	}
}

func (r userRepository) GetAll() ([]models.Result, error) {
	//TODO implement me
	panic("implement me")
}
