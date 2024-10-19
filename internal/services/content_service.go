package services

import (
	"context"
	"dune-imperium-service/internal/dto/api"
	"dune-imperium-service/internal/models"
	"dune-imperium-service/internal/repositories"
	"errors"
	"github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

type ContentService struct {
	logger         *logrus.Logger
	contentRepo    *repositories.ContentRepository
	storageService *FileStorageService
}

func NewContentService(
	logger *logrus.Logger,
	contentRepo *repositories.ContentRepository,
	storageService *FileStorageService,
) *ContentService {
	return &ContentService{
		logger:         logger,
		contentRepo:    contentRepo,
		storageService: storageService,
	}
}

func (s *ContentService) Create(ctx context.Context, request *api.ContentCreateRequest) (string, error) {
	id, err := s.generateID(request.Name)
	if err != nil {
		return "", err
	}

	imageUrl, err := s.storageService.UploadFile(ctx, "content", id, request.Image)
	if err != nil {
		return "", err
	}

	content := &models.GameContent{
		ID:       id,
		Name:     request.Name,
		Type:     request.Type,
		ImageURL: imageUrl,
	}

	contentId, err := s.contentRepo.Save(ctx, content)
	if err != nil {
		return "", err
	}

	return contentId, nil
}

func (s *ContentService) GetById(ctx context.Context, id string) (*models.GameContent, error) {
	return s.contentRepo.FindById(ctx, id)
}
func (s *ContentService) GetByType(ctx context.Context, contentType models.ContentType) ([]*models.GameContent, error) {
	return s.contentRepo.FindByType(ctx, contentType)
}

func (s *ContentService) generateID(name string) (string, error) {
	if name == "" {
		return "", errors.New("cannot generate id for empty name")
	}

	name = regexp.MustCompile(`[^a-zA-Z0-9\s]`).ReplaceAllString(name, "")
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ToLower(name)

	return name, nil
}
