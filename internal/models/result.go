package models

import "time"

type PlayerResult struct {
	PlayerId string `bson:"player_id"`
	Leader   string `bson:"leader"`
	Rank     int    `bson:"rank"`
	Points   int    `bson:"points"`
}

type Result struct {
	GameID        string         `bson:"_id"`
	GameMode      string         `bson:"game_mode"`
	Date          time.Time      `bson:"date"`
	ScreenshotURL string         `bson:"screenshot_url"`
	Results       []PlayerResult `bson:"results"`
}
