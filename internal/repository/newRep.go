package repository

import (
	"builder_service/internal/model"
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	// ErrNotFound means entity is not found in repository
	ErrNotFound = errors.New("not found")
)

// Postgres struct for Pool
type Postgres struct {
	Pool *pgxpool.Pool
}

// NewRepository set new repository for mongo and postgres
func NewRepository(db interface{}, dbType string) Building {
	return &Postgres{Pool: db.(*pgxpool.Pool)}
}

// Building used for structuring, function for working with records
type Building interface {
	InsertRecord(c context.Context, building *model.Building) error
	SelectRecord(c context.Context, idPlanet uuid.UUID, name int) (*model.Building, error)
	SelectAllRecords(c context.Context, idPlanet uuid.UUID) ([]*model.Building, error)
	//	UpdateRecord(c context.Context, building *model.Building) error
	DeleteRecord(c context.Context, idPlanet uuid.UUID, name int) error
}

// ShipYard used for structuring, function for working with records
type ShipYard interface {
	InsertRecord(c context.Context, ship *model.Ship) error
	SelectRecord(c context.Context, idPlanet uuid.UUID, name int) (*model.Ship, error)
	SelectAllRecords(c context.Context, idPlanet uuid.UUID) ([]*model.Ship, error)
	UpdateRecord(c context.Context, ship *model.Ship) error
	DeleteRecord(c context.Context, idPlanet uuid.UUID, name int) error
}

// Laboratory used for structuring, function for working with records
type Laboratory interface {
	InsertRecord(c context.Context, research *model.Research) error
	SelectRecord(c context.Context, idPlanet uuid.UUID, name int) (*model.Research, error)
	SelectAllRecords(c context.Context, idPlanet uuid.UUID) ([]*model.Research, error)
	UpdateRecord(c context.Context, research *model.Research) error
	DeleteRecord(c context.Context, idPlanet uuid.UUID, name int) error
}
