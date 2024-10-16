package mappers

import (
	"dune-imperium-service/internal/dto/api"
	"dune-imperium-service/internal/models"
)

func ToPlayerResponse(player *models.Player) *api.PlayerResponse {
	return &api.PlayerResponse{
		ID:        player.ID,
		Nickname:  player.Nickname,
		AvatarURL: player.AvatarURL,
	}
}
