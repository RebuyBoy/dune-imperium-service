package services

import (
	"hs-fake-api/internal/model"
	"log"
	"math/rand"
)

type DataGenerator struct {
	l *log.Logger
}

func NewDataGenerator(l *log.Logger) *DataGenerator {
	return &DataGenerator{l}
}

func (d *DataGenerator) Generate(startID, density, durationSeconds int) model.FakeHSTournament {
	var tournaments []*model.FakeTournament

	for i := 0; i < durationSeconds*density; i++ {
		tournaments = append(tournaments, NewFakeTournament(startID))
		startID++
	}

	return model.FakeHSTournament{Tournaments: tournaments}
}

const (
	Winamax    model.Room = "WINAMAX"
	PokerStars model.Room = "POKER_STARS"
	GG         model.Room = "GG_POKER"
)

const (
	COM model.Reservation = "COM"
)

const (
	Turbo   model.Speed = "TURBO"
	Regular model.Speed = "REGULAR"
)

var rooms = []model.Room{Winamax, PokerStars, GG}
var speeds = []model.Speed{Turbo, Regular}
var stakes = []model.Stake{500, 250, 100, 50, 25, 10, 5}

func NewFakeTournament(id int) *model.FakeTournament {
	room := rooms[rand.Intn(len(rooms))]
	speed := speeds[rand.Intn(len(speeds))]
	stake := stakes[rand.Intn(len(stakes))]
	playerId := getPlayerID(room, speed)

	return &model.FakeTournament{
		ID:          id,
		Room:        room,
		Reservation: COM,
		Speed:       speed,
		Stake:       stake,
		PlayerID:    playerId,
		//Date:        time.Now(),
	}
}

func getPlayerID(room model.Room, speed model.Speed) int {
	switch room {
	case Winamax:
		if speed == Regular {
			return rand.Intn(15)
		}
		return 101 + rand.Intn(30)
	case PokerStars:
		if speed == Regular {
			return 201 + rand.Intn(15)
		}
		return 301 + rand.Intn(30)
	case GG:
		if speed == Regular {
			return 401 + rand.Intn(15)
		}
		return 501 + rand.Intn(30)
	default:
		return 0
	}
}
