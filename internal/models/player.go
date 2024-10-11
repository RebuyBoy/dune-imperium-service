package models

import (
	"time"
)

type Player struct {
	ID           string    `db:"id"`
	Nickname     string    `db:"nickname"`
	Email        string    `db:"email"`
	Avatar       []byte    `db:"avatar"`
	RegisteredAt time.Time `db:"registered_at"`
}
