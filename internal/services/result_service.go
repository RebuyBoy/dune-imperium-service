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
)

type ResultService interface {
	GetAll(ctx context.Context) ([]models.Result, error)
	Save(ctx context.Context, saveRequest api.ResultSaveRequest) error
}

type resultService struct {
	logger      *logrus.Logger
	resultRepo  *repositories.ResultRepository
	minioClient *minio.Client
}

func NewResultService(logger *logrus.Logger, resultRepo *repositories.ResultRepository, minioClient *minio.Client) ResultService {
	return &resultService{
		logger:      logger,
		resultRepo:  resultRepo,
		minioClient: minioClient}
}

func (s *resultService) GetAll(ctx context.Context) ([]models.Result, error) {
	results, err := s.resultRepo.GetAll(ctx)
	if err != nil {
		s.logger.Error("Error fetching results: ", err)
		return nil, err
	}
	return results, nil
}

func (s *resultService) Save(ctx context.Context, saveRequest api.ResultSaveRequest) error {
	gameId := uuid.New().String()
	var screenshotURL string
	var err error
	if saveRequest.Screenshot != nil {
		screenshotURL, err = s.uploadScreenshot(gameId, saveRequest.Screenshot)
		if err != nil {
			s.logger.Error("Error uploading screenshot: ", err)
			return err
		}
	}

	result := models.Result{
		GameId:        gameId,
		GameMode:      saveRequest.GameMode,
		Date:          saveRequest.Date,
		ScreenshotURL: screenshotURL,
		Results:       mapPlayerResults(saveRequest.Results),
	}
	err = s.resultRepo.Save(ctx, result)
	if err != nil {
		s.logger.Error("Error saving result: ", err)
		return err
	}
	return nil
}

func mapPlayerResults(requests []api.PlayerResultRequest) []models.PlayerResult {
	var results []models.PlayerResult
	for _, req := range requests {
		results = append(results, models.PlayerResult{
			PlayerId: req.PlayerId,
			Leader:   req.Leader,
			Rank:     req.Rank,
			Points:   req.Points,
		})
	}
	return results
}

func (s *resultService) uploadScreenshot(GameId string, screenshot *api.Screenshot) (string, error) {
	bucketName := "screenshots"

	ext := filepath.Ext(screenshot.Filename)
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}
	objectName := fmt.Sprintf("%s%s", GameId, ext)

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

	_, err = s.minioClient.PutObject(context.Background(), bucketName, objectName, screenshot.File, screenshot.Size, minio.PutObjectOptions{
		ContentType: mimeType,
	})
	if err != nil {
		return "", err
	}

	avatarURL := fmt.Sprintf("%s/%s/%s", s.minioClient.EndpointURL().String(), bucketName, objectName)
	//TODO endpoint to abstract /storage extract to minioService

	return avatarURL, nil
}
