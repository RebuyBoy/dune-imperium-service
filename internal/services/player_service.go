package services

import (
	"context"
	"dune-imperium-service/internal/dto/api"
	"dune-imperium-service/internal/models"
	"dune-imperium-service/internal/repositories"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
	"mime"
	"path/filepath"
	"time"
)

type PlayerService interface {
	Create(ctx context.Context, request api.PlayerCreateRequest) error
	GetNames(ctx context.Context) ([]string, error)
}

type playerService struct {
	logger      *logrus.Logger
	playerRepo  repositories.PlayerRepository
	minioClient *minio.Client
}

func NewPlayerService(logger *logrus.Logger, playerRepo repositories.PlayerRepository, minioClient *minio.Client) PlayerService {
	return &playerService{
		logger:      logger,
		playerRepo:  playerRepo,
		minioClient: minioClient,
	}
}

func (s *playerService) Create(ctx context.Context, request api.PlayerCreateRequest) error {

	player := &models.Player{
		ID:           uuid.New().String(),
		Nickname:     request.Nickname,
		RegisteredAt: time.Now(),
	}

	time.Sleep(5 * time.Second)

	avatarURL, err := s.uploadAvatar(player.ID, request.Avatar)
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

func (s *playerService) GetNames(ctx context.Context) ([]string, error) {
	names, err := s.playerRepo.GetNames(ctx)
	if err != nil {
		s.logger.Error("Error retrieving player names: ", err)
		return nil, err
	}
	return names, nil
}

func (s *playerService) uploadAvatar(playerID string, avatar api.Avatar) (string, error) {
	bucketName := "avatars"

	ext := filepath.Ext(avatar.Filename)
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}
	objectName := fmt.Sprintf("%s%s", playerID, ext)

	exists, err := s.minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		return "", err
	}
	if !exists {
		err = s.minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return "", err
		}
	}

	_, err = s.minioClient.PutObject(context.Background(), bucketName, objectName, avatar.File, avatar.Size, minio.PutObjectOptions{
		ContentType: mimeType,
	})
	if err != nil {
		return "", err
	}

	avatarURL := fmt.Sprintf("%s/%s/%s", s.minioClient.EndpointURL().String(), bucketName, objectName)
	//TODO endpoint to abstract /storage

	return avatarURL, nil
}
