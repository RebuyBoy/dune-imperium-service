package services

import (
	"context"
	"dune-imperium-service/internal/dto/api"
	"dune-imperium-service/internal/models"
	"dune-imperium-service/internal/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"time"
)

type PlayerService struct {
	logger         *logrus.Logger
	playerRepo     repositories.PlayerRepository
	storageService *FileStorageService
}

func NewPlayerService(logger *logrus.Logger, playerRepo repositories.PlayerRepository, storageService *FileStorageService) *PlayerService {
	return &PlayerService{
		logger:         logger,
		playerRepo:     playerRepo,
		storageService: storageService,
	}
}

func (s *PlayerService) Create(ctx context.Context, request api.PlayerCreateRequest) error {
	player := &models.Player{
		ID:           uuid.New().String(),
		Nickname:     request.Nickname,
		RegisteredAt: time.Now(),
	}

	time.Sleep(5 * time.Second)

	avatarURL, err := s.uploadAvatar(ctx, player.ID, request.Avatar)
	if err != nil {
		s.logger.Error("Error uploading avatar to MinIO: ", err)
		return err
	}

	player.AvatarURL = avatarURL

	err = s.playerRepo.Save(ctx, player)
	if err != nil {
		s.logger.Error("Error saving player to MongoDB: ", err)
		return err
	}

	return nil
}

func (s *PlayerService) GetNames(ctx context.Context) ([]string, error) {
	names, err := s.playerRepo.GetNames(ctx)
	if err != nil {
		s.logger.Error("Error retrieving player names: ", err)
		return nil, err
	}
	return names, nil
}

func (s *PlayerService) uploadAvatar(ctx context.Context, playerID string, avatar *models.FileData) (string, error) {
	return s.storageService.UploadFile(ctx, "avatars", playerID, avatar)
}
