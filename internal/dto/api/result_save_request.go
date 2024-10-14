package api

import (
	"io"
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
	Screenshot *Screenshot           `json:"-"`
	Results    []PlayerResultRequest `json:"results"`
}

type Screenshot struct {
	File     io.Reader
	Size     int64
	Filename string
}
