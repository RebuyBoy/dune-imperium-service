package api

import (
	"dune-imperium-service/internal/models"
)

type PlayerCreateRequest struct {
	Nickname string           `json:"nickname"`
	Avatar   *models.FileData `json:"-"`
}
