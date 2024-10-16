package models

import (
	"time"
)

type Player struct {
	ID           string    `bson:"_id"`
	Nickname     string    `bson:"nickname"`
	AvatarURL    string    `bson:"avatar_url"`
	RegisteredAt time.Time `bson:"registered_at"`
}
