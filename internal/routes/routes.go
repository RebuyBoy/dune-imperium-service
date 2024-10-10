package routes

import (
	"github.com/gofiber/fiber/v2"
	"hs-fake-api/internal/service"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type Handler struct {
	l *log.Logger
	a fiber.Router
}

func NewHandler(
	logger *log.Logger,
	router fiber.Router,
	dataGenerator *service.DataGenerator) *Handler {
	return &Handler{logger, router, dataGenerator}
}

func (h *Handler) RegisterRoutes() {
	h.a.Get("/tournaments", h.tournamentsHandler)
	h.a.Post("/observer/stats", h.liveStatsHandler)
}

type LiveStats struct {
	Live LiveData `json:"live"`
	Max  MaxData  `json:"max"`
}

type LiveData struct {
	Players int       `json:"players"`
	Table   TableData `json:"table"`
}

type MaxData struct {
	Players int       `json:"players"`
	Table   TableData `json:"table"`
}

type TableData struct {
	Count   int `json:"count"`
	Average int `json:"average"`
}

func (h *Handler) liveStatsHandler(c *fiber.Ctx) error {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	livePlayers := rng.Intn(100) + 1
	liveTableCount := rng.Intn(200) + 50
	liveTableAverage := liveTableCount / 2

	maxPlayers := 200
	maxTableCount := 400
	maxTableAverage := maxTableCount / 2

	return c.JSON(LiveStats{
		Live: LiveData{
			Players: livePlayers,
			Table: TableData{
				Count:   liveTableCount,
				Average: liveTableAverage,
			},
		},
		Max: MaxData{
			Players: maxPlayers,
			Table: TableData{
				Count:   maxTableCount,
				Average: maxTableAverage,
			},
		},
	})
}

func (h *Handler) tournamentsHandler(c *fiber.Ctx) error {
	startIDStr := c.Query("start_id")
	densityStr := c.Query("density")
	timeStr := c.Query("seconds")

	startID, err := strconv.Atoi(startIDStr)
	if err != nil || startID <= 0 {
		h.l.Printf("Invalid start id parameter: %s", startIDStr)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid start id parameter")
	}

	density, err := strconv.Atoi(densityStr)
	if err != nil || density <= 0 {
		h.l.Printf("Invalid density parameter: %s", densityStr)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid density parameter")
	}

	durationSeconds, err := strconv.Atoi(timeStr)
	if err != nil || durationSeconds <= 0 {
		h.l.Printf("Invalid time parameter: %s", timeStr)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid seconds parameter")
	}

	data := h.s.Generate(startID, density, durationSeconds)

	return c.JSON(data)
}
