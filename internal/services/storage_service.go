package services

import (
	"context"
	"dune-imperium-service/internal/models"
	"fmt"
	"mime"
	"path/filepath"

	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
)

type FileStorageService struct {
	logger      *logrus.Logger
	minioClient *minio.Client
	baseURL     string
}

func NewFileStorageService(logger *logrus.Logger, minioClient *minio.Client, baseURL string) *FileStorageService {
	return &FileStorageService{
		logger:      logger,
		minioClient: minioClient,
		baseURL:     baseURL,
	}
}

func (s *FileStorageService) UploadFile(ctx context.Context, bucketName, objectID string, file *models.FileData) (string, error) {
	ext := filepath.Ext(file.Filename)
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}
	objectName := fmt.Sprintf("%s%s", objectID, ext)

	exists, err := s.minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return "", fmt.Errorf("failed to check bucket existence: %w", err)
	}
	if !exists {
		err = s.minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return "", fmt.Errorf("failed to create bucket: %w", err)
		}
	}

	_, err = s.minioClient.PutObject(ctx, bucketName, objectName, file.Content, file.Size, minio.PutObjectOptions{
		ContentType: mimeType,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	fileURL := fmt.Sprintf("%s/%s/%s", s.baseURL, bucketName, objectName)
	return fileURL, nil
}
