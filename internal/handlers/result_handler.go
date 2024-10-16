package handlers

import (
	"dune-imperium-service/internal/dto/api"
	"dune-imperium-service/internal/models"
	"dune-imperium-service/internal/services"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"time"
)

type ResultHandler struct {
	logger        *logrus.Logger
	resultService *services.ResultService
}

func NewResultHandler(logger *logrus.Logger, resultService *services.ResultService) *ResultHandler {
	return &ResultHandler{logger: logger, resultService: resultService}
}

func (h *ResultHandler) GetAll(c *fiber.Ctx) error {
	//TODO pagination
	//TODO  playerId to nickname, join query?
	results, err := h.resultService.GetAll(c.Context())
	if err != nil {
		return h.handleError(c, fiber.StatusInternalServerError, "Failed to fetch results", err)
	}
	return c.JSON(results)
}

func (h *ResultHandler) Save(c *fiber.Ctx) error {
	var request api.ResultSaveRequest

	request, err := h.parseSaveRequest(c)
	if err != nil {
		return h.handleError(c, fiber.StatusBadRequest, "Invalid request", err)
	}

	screenshot, err := h.processScreenshot(c)
	if err != nil {
		return h.handleError(c, fiber.StatusBadRequest, "Failed to process screenshot", err)
	}
	request.Screenshot = screenshot

	err = h.resultService.Save(c.Context(), request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "failed to save results",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Result saved successfully"})
}

func (h *ResultHandler) parseSaveRequest(c *fiber.Ctx) (api.ResultSaveRequest, error) {
	var request api.ResultSaveRequest

	form, err := c.MultipartForm()
	if err != nil {
		return request, fmt.Errorf("failed to parse multipart form: %w", err)
	}

	if gameModes := form.Value["game_mode"]; len(gameModes) > 0 {
		request.GameMode = gameModes[0]
	}

	if dates := form.Value["date"]; len(dates) > 0 {
		parsedDate, err := time.Parse(time.RFC3339, dates[0])
		if err != nil {
			return request, fmt.Errorf("invalid date format: %w", err)
		}
		request.Date = parsedDate
	}

	if resultsJSON := form.Value["results"]; len(resultsJSON) > 0 {
		if err := json.Unmarshal([]byte(resultsJSON[0]), &request.Results); err != nil {
			return request, fmt.Errorf("failed to parse results: %w", err)
		}
	}

	if request.GameMode == "" {
		return request, fmt.Errorf("game_mode is required")
	}
	if request.Date.IsZero() {
		return request, fmt.Errorf("date is required")
	}

	if len(request.Results) < 2 || len(request.Results) > 4 {
		return request, fmt.Errorf("should be between 2 and 4 players")
	}

	return request, nil
}

func (h *ResultHandler) processScreenshot(c *fiber.Ctx) (*models.FileData, error) {
	fileHeader, err := c.FormFile("screenshot")
	if err != nil {
		return nil, err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return &models.FileData{
		Content:  file,
		Size:     fileHeader.Size,
		Filename: fileHeader.Filename,
	}, nil
}

func (h *ResultHandler) handleError(c *fiber.Ctx, status int, message string, err error) error {
	h.logger.Error(message+": ", err)
	return c.Status(status).JSON(fiber.Map{"error": message})
}
