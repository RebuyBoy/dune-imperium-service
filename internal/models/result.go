package models

import "time"

type PlayerResult struct {
	PlayerId string
	Leader   string
	Rank     int
	Points   int
}

type Result struct {
	GameId        string
	GameMode      string
	Date          time.Time
	ScreenshotURL string
	Results       []PlayerResult
}
