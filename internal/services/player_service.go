package services

import (
	"dune-imperium-service/internal/models"
	"dune-imperium-service/internal/repositories"
	"github.com/sirupsen/logrus"
)

type PlayerService interface {
	GetAll() []models.Result
}

type userService struct {
	logger     *logrus.Logger
	playerRepo repositories.PlayerRepository
}

func NewUserService(logger *logrus.Logger, userRepo repositories.PlayerRepository) PlayerService {
	return &userService{logger, userRepo}
}

func (gs userService) GetAll() []models.Result {
	return nil
}
