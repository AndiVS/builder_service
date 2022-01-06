package repository

import (
	"builder_service/internal/model"
	"context"
	"github.com/google/uuid"

	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
)

// InsertRecord function for inserting item from a table
func (repos *Postgres) InsertRecord(c context.Context, building *model.Building) error {
	row := repos.Pool.QueryRow(c,
		"INSERT INTO buildings (id_user, id_planet, type, name, time_start, time_need, update, progress) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING progress",
		building.IDUser, building.IDPlanet, building.Type, building.Name, building.TimeStart, building.TimeNeed, building.Update, building.Progress)

	err := row.Scan(&building.Progress)
	if err != nil {
		log.Errorf("Unable to INSERT: %v", err)
		return err
	}

	return err
}

// SelectRecord function for selecting item from a table
func (repos *Postgres) SelectRecord(c context.Context, idPlanet uuid.UUID, name int) (*model.Building, error) {
	var building model.Building
	row := repos.Pool.QueryRow(c,
		"SELECT id_user, id_planet, type, name, time_start, time_need, update, progress  FROM buildings WHERE id_planet = $1 AND name = $2", idPlanet, name)

	err := row.Scan(&building.IDUser, &building.IDPlanet, &building.Type, &building.Name, &building.TimeStart, &building.TimeNeed, &building.Update, &building.Progress)
	if err != nil {
		log.Errorf("Unable to SELECT: %v", err)
		return &building, err
	}

	log.Printf("sec")

	return &building, err
}

// SelectAllRecords function for selecting items from a table
func (repos *Postgres) SelectAllRecords(c context.Context, idPlanet uuid.UUID) ([]*model.Building, error) {
	var buildings []*model.Building

	row, err := repos.Pool.Query(c,
		"SELECT id_user, id_planet, type, name, time_start, time_need, update, progress  FROM buildings WHERE id_planet = $1", idPlanet)

	for row.Next() {
		var building model.Building
		err := row.Scan(&building.IDUser, &building.IDPlanet, &building.Type, &building.Name, &building.TimeStart, &building.TimeNeed, &building.Update, &building.Progress)
		if err == pgx.ErrNoRows {
			return buildings, err
		}
		buildings = append(buildings, &building)
	}

	return buildings, err
}

/*
// UpdateRecord function for updating item from a table
func (repos *Postgres) UpdateRecord(c context.Context, building *model.Building) error {

	_, err := repos.Pool.Exec(c,
		"UPDATE buildings SET is_admin = $2 WHERE username = $1", username, isAdmin)
	if err != nil {
		log.Errorf("Failed updating data in db: %s\n", err)
		return err
	}

	return nil
}*/

// DeleteRecord function for deleting item from a table
func (repos *Postgres) DeleteRecord(c context.Context, idPlanet uuid.UUID, name int) error {
	ct, err := repos.Pool.Exec(c, "DELETE FROM buildings WHERE id_planet = $1 AND name = $2", idPlanet, name)

	if err != nil {
		return err
	} else if ct.RowsAffected() == 0 {
		log.Errorf("Not found : %s\n", err)
		return ErrNotFound
	}

	return nil
}
