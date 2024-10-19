package validation

import (
	"dune-imperium-service/internal/models"
	"fmt"
)

func ContentType(typeStr string) (models.ContentType, error) {
	contentType := models.ContentType(typeStr)
	switch contentType {
	case models.Leader, models.Tech, models.Intrigue, models.Card:
		return contentType, nil
	default:
		return "", fmt.Errorf("invalid content type: %s", typeStr)
	}
}
