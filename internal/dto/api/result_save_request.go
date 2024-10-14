package api

import (
	"dune-imperium-service/internal/models"
	"time"
)

type PlayerResultRequest struct {
	PlayerId string `json:"player_id"`
	Leader   string `json:"leader"`
	Rank     int    `json:"rank"`
	Points   int    `json:"points"`
}

type ResultSaveRequest struct {
	GameMode   string                `json:"game_mode"`
	Date       time.Time             `json:"date"`
	Screenshot *models.FileData      `json:"-"`
	Results    []PlayerResultRequest `json:"results"`
}
