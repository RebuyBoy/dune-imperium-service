package handlers

import (
	"dune-imperium-service/internal/dto/api"
	"dune-imperium-service/internal/models"
	"dune-imperium-service/internal/services"
	"dune-imperium-service/internal/validation"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"mime/multipart"
)

type ContentHandler struct {
	logger         *logrus.Logger
	contentService *services.ContentService
}

func NewContentHandler(
	logger *logrus.Logger,
	contentService *services.ContentService,
) *ContentHandler {
	return &ContentHandler{
		logger:         logger,
		contentService: contentService,
	}
}

func (h *ContentHandler) Create(c *fiber.Ctx) error {
	request, err := h.parseContentCreateRequest(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request", "details": err.Error()})
	}
	id, err := h.contentService.Create(c.Context(), request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}

func (h *ContentHandler) GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "id is required"})
	}

	content, err := h.contentService.GetById(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get content"})
	}

	return c.Status(fiber.StatusOK).JSON(content)
}

func (h *ContentHandler) GetByType(c *fiber.Ctx) error {
	typeStr := c.Params("type")
	if typeStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "type is required"})
	}

	contentType, err := validation.ContentType(typeStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid content type"})
	}
	content, err := h.contentService.GetByType(c.Context(), contentType)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get content"})
	}
	return c.Status(fiber.StatusOK).JSON(content)
}

func (h *ContentHandler) parseContentCreateRequest(c *fiber.Ctx) (*api.ContentCreateRequest, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, fmt.Errorf("failed to parse multipart form: %w", err)
	}

	name := h.getFormValue(form, "name")
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}

	typeStr := h.getFormValue(form, "type")
	if typeStr == "" {
		return nil, fmt.Errorf("type is required")
	}

	contentType, err := validation.ContentType(typeStr)
	if err != nil {
		return nil, fmt.Errorf("invalid content type: %w", err)
	}

	image, err := h.processImage(c)
	if err != nil {
		return nil, fmt.Errorf("failed to process image: %w", err)
	}

	return &api.ContentCreateRequest{
		Name:  name,
		Type:  contentType,
		Image: image,
	}, nil
}

func (h *ContentHandler) getFormValue(form *multipart.Form, key string) string {
	if values := form.Value[key]; len(values) > 0 {
		return values[0]
	}
	return ""
}

func (h *ContentHandler) processImage(c *fiber.Ctx) (*models.FileData, error) {
	fileHeader, err := c.FormFile("image")
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
