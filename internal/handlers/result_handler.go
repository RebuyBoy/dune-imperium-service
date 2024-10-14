package handlers

import (
	"dune-imperium-service/internal/dto/api"
	"dune-imperium-service/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ResultHandler struct {
	logger        *logrus.Logger
	resultService services.ResultService
}

func NewResultHandler(logger *logrus.Logger, resultService services.ResultService) *ResultHandler {
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
			"error": "failed to save results",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Result saved successfully"})
}

func (h *ResultHandler) parseSaveRequest(c *fiber.Ctx) (api.ResultSaveRequest, error) {
	var request api.ResultSaveRequest
	if err := c.BodyParser(&request); err != nil {
		return request, err
	}
	return request, nil
}

func (h *ResultHandler) processScreenshot(c *fiber.Ctx) (*api.Screenshot, error) {
	fileHeader, err := c.FormFile("screenshot")
	if err != nil {
		return nil, err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return &api.Screenshot{
		File:     file,
		Size:     fileHeader.Size,
		Filename: fileHeader.Filename,
	}, nil
}

func (h *ResultHandler) handleError(c *fiber.Ctx, status int, message string, err error) error {
	h.logger.Error(message+": ", err)
	return c.Status(status).JSON(fiber.Map{"error": message})
}
