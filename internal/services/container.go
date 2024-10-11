package services

import (
	"dune-imperium-service/internal/repositories"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Container struct {
	ResultService ResultService
	PlayerService PlayerService
}

func NewServiceContainer(logger *logrus.Logger, dbClient *mongo.Client) *Container {
	resultRepo := repositories.NewResultRepository(dbClient)
	resultService := NewResultService(logger, resultRepo)

	playerRepo := repositories.NewPlayerRepository(dbClient)
	playerService := NewUserService(logger, playerRepo)

	return &Container{
		ResultService: resultService,
		PlayerService: playerService,
	}
}
