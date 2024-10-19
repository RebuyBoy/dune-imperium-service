package api

import "dune-imperium-service/internal/models"

type ContentCreateRequest struct {
	Name  string             `json:"name"`
	Type  models.ContentType `json:"type"`
	Image *models.FileData   `json:"-"`
}
