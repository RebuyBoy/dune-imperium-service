package services

import (
	"dune-imperium-service/internal/repositories"
	"github.com/sirupsen/logrus"
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
