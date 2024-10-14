package services

import (
	"dune-imperium-service/internal/repositories"
	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Container struct {
	ResultService ResultService
	PlayerService PlayerService
}

type ServiceDependencies struct {
	Logger      *logrus.Logger
	MongoClient *mongo.Client
	MinioClient *minio.Client
}

func NewServiceContainer(deps ServiceDependencies) *Container {
	resultRepo := repositories.NewResultRepository(deps.MongoClient)
	resultService := NewResultService(deps.Logger, resultRepo, deps.MinioClient)

	playerRepo := repositories.NewPlayerRepository(deps.MongoClient)
	playerService := NewPlayerService(deps.Logger, playerRepo, deps.MinioClient)

	return &Container{
		ResultService: resultService,
		PlayerService: playerService,
	}
}
