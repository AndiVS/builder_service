package model

import (
	"github.com/google/uuid"
	"time"
)

type Building struct {
	UserID       uuid.UUID     `json:"user_id"`
	PlanetID     uuid.UUID     `json:"planet_id"`
	BuildingType int           `json:"building_type"`
	BuildingID   int           `json:"building_id"`
	Destroy      bool          `json:"destroy"`
	Start        time.Time     `json:"start"`
	Duration     time.Duration `json:"duration"`
	Status       int           `json:"status"`
}

type Ship struct {
	UserID       uuid.UUID     `json:"user_id"`
	PlanetID     uuid.UUID     `json:"planet_id"`
	ShipYardType int           `json:"ship_yard_type"`
	ShipYardID   int           `json:"ship_yard_id"`
	Destroy      bool          `json:"destroy"`
	Start        time.Time     `json:"start"`
	Duration     time.Duration `json:"duration"`
	Status       int           `json:"status"`
}

type Research struct {
	UserID     uuid.UUID     `json:"user_id"`
	PlanetID   uuid.UUID     `json:"planet_id"`
	ResearchID int           `json:"research_id"`
	Start      time.Time     `json:"start"`
	Duration   time.Duration `json:"duration"`
	Status     int           `json:"status"`
}
