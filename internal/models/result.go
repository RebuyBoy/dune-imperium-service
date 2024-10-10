package models

import "time"

type PlayerResult struct {
	PlayerId string
	Race     string
	Leader   string
	Rank     int
	Points   int
}

type Result struct {
	GameId   int
	GameMode string
	PlayedAt time.Time
	Results  []PlayerResult
}
