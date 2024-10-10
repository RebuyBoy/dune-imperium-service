package services

import (
	"dune-imperium-service/internal/models"
	"dune-imperium-service/internal/repositories"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	GetAll() []models.Result
}

type userService struct {
	logger   *logrus.Logger
	userRepo repositories.UserRepository
}

func NewUserService(logger *logrus.Logger, userRepo repositories.UserRepository) UserService {
	return &userService{logger, userRepo}
}

func (gs userService) GetAll() []models.Result {
	return nil
}
