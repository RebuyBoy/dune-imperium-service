package services

import (
	"context"
	"dune-imperium-service/internal/dto/api"
	"dune-imperium-service/internal/models"
	"dune-imperium-service/internal/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ResultService struct {
	logger         *logrus.Logger
	resultRepo     *repositories.ResultRepository
	storageService *FileStorageService
}

func NewResultService(logger *logrus.Logger, resultRepo *repositories.ResultRepository, fileStorageService *FileStorageService) *ResultService {
	return &ResultService{
		logger:         logger,
		resultRepo:     resultRepo,
		storageService: fileStorageService,
	}
}

func (s *ResultService) GetAll(ctx context.Context) ([]models.Result, error) {
	results, err := s.resultRepo.GetAll(ctx)
	if err != nil {
		s.logger.Error("Error fetching results: ", err)
		return nil, err
	}
	return results, nil
}

func (s *ResultService) Save(ctx context.Context, saveRequest api.ResultSaveRequest) error {
	gameId := uuid.New().String()
	var screenshotURL string
	var err error
	if saveRequest.Screenshot != nil {
		screenshotURL, err = s.uploadScreenshot(ctx, gameId, saveRequest.Screenshot)
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

func (s *ResultService) uploadScreenshot(ctx context.Context, gameID string, screenshot *models.FileData) (string, error) {
	return s.storageService.UploadFile(ctx, "screenshots", gameID, screenshot)
}
