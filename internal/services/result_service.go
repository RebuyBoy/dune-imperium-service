package services

import (
	"dune-imperium-service/internal/models"
	"dune-imperium-service/internal/repositories"
	"github.com/sirupsen/logrus"
)

type ResultService interface {
	GetAll() []models.Result
}

type resultService struct {
	logger   *logrus.Logger
	gameRepo repositories.ResultRepository
}

func NewResultService(logger *logrus.Logger, gameRepo repositories.ResultRepository) ResultService {
	return &resultService{logger, gameRepo}
}

func (gs resultService) GetAll() []models.Result {
	return nil
}
