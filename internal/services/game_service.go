package services

import (
	"dune-imperium-service/internal/models"
	"dune-imperium-service/internal/repositories"
	"github.com/sirupsen/logrus"
)

type GameService interface {
	GetAll() []models.Result
}

type gameService struct {
	logger   *logrus.Logger
	gameRepo repositories.GameRepository
}

func NewGameService(logger *logrus.Logger, gameRepo repositories.GameRepository) GameService {
	return &gameService{logger, gameRepo}
}

func (gs gameService) GetAll() []models.Result {
	return nil
}
