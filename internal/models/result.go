package models

import "time"

type PlayerResult struct {
	PlayerId string
	Race     string
	GameMode string
	Leader   string
	Rank     int
	Points   int
}

type Result struct {
	GameId     int
	StartedAt  time.Time
	FinishedAt time.Time
	Results    []PlayerResult
}
