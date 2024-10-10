package services

import (
	"dune-imperium-service/internal/repositories"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Container struct {
	GameService UserService
	UserService UserService
}

func NewServiceContainer(logger *logrus.Logger, dbClient *mongo.Client) *Container {
	gameRepo := repositories.NewResultRepository(dbClient)
	gameService := NewResultService(logger, gameRepo)

	userRepo := repositories.NewUserRepository(dbClient)
	userService := NewUserService(logger, userRepo)

	return &Container{
		GameService: gameService,
		UserService: userService,
	}
}
