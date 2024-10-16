package services

import (
	"context"
	"dune-imperium-service/internal/dto/api"
	"dune-imperium-service/internal/models"
	"dune-imperium-service/internal/repositories"
	"errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"time"
)

type PlayerService struct {
	logger         *logrus.Logger
	playerRepo     *repositories.PlayerRepository
	storageService *FileStorageService
}

func NewPlayerService(
	logger *logrus.Logger,
	playerRepo *repositories.PlayerRepository,
	storageService *FileStorageService,
) *PlayerService {
	return &PlayerService{
		logger:         logger,
		playerRepo:     playerRepo,
		storageService: storageService,
	}
}

func (s *PlayerService) Create(ctx context.Context, request api.PlayerCreateRequest) (*models.Player, error) {

	isExists, err := s.playerRepo.IsNicknameExists(ctx, request.Nickname)
	if err != nil {
		s.logger.Error("Error checking nickname uniqueness: ", err)
		return nil, err
	}
	if isExists {
		return nil, errors.New("nickname already exists")
	}

	player := &models.Player{
		ID:           uuid.New().String(),
		Nickname:     request.Nickname,
		RegisteredAt: time.Now(),
	}

	avatarURL, err := s.uploadAvatar(ctx, player.ID, request.Avatar)
	if err != nil {
		s.logger.Error("Error uploading avatar to MinIO: ", err)
		return nil, err
	}

	player.AvatarURL = avatarURL

	err = s.playerRepo.Save(ctx, player)
	if err != nil {
		s.logger.Error("Error saving player to MongoDB: ", err)
		return nil, err
	}

	return player, nil
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

func (s *PlayerService) GetById(ctx context.Context, playerID string) (*models.Player, error) {
	player, err := s.playerRepo.GetById(ctx, playerID)
	if err != nil {
		s.logger.Error("Error retrieving player from MongoDB: ", err)
		return nil, err
	}
	return player, nil
}
