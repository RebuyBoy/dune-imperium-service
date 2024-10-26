package models

import "time"

type PlayerResult struct {
	PlayerId string `bson:"player_id" json:"player_id"`
	Leader   string `bson:"leader" json:"leader"`
	Rank     int    `bson:"rank" json:"rank"`
	Points   int    `bson:"points" json:"points"`
}

type GameResult struct {
	ID            string         `bson:"_id" json:"id"`
	GameMode      string         `bson:"game_mode" json:"game_mode"`
	Date          time.Time      `bson:"date" json:"date"`
	UploadedAt    time.Time      `bson:"uploaded_at" json:"uploaded_at"`
	UploadedBy    string         `bson:"uploaded_by" json:"uploaded_by"`
	ScreenshotURL string         `bson:"screenshot_url" json:"screenshot_url"`
	Results       []PlayerResult `bson:"results" json:"results"`
}
