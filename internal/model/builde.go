package model

import (
	"github.com/google/uuid"
)

type Building struct {
	IDUser    uuid.UUID `json:"id_user"`
	IDPlanet  uuid.UUID `json:"id_planet"`
	Type      int       `json:"type"`
	Name      int       `json:"name"`
	TimeStart int       `json:"time_start"`
	TimeNeed  int       `json:"time_need"`
	Update    bool      `json:"update"`
	Progress  int       `json:"progress"`
}

type Ship struct {
	IDUser    uuid.UUID `json:"id_user"`
	IDPlanet  uuid.UUID `json:"id_planet"`
	Type      int       `json:"RequestType"`
	Name      int       `json:"Name"`
	TimeStart int       `json:"time_start"`
	TimeNeed  int       `json:"time_need"`
	Progress  int       `json:"progress"`
}

type Research struct {
	IDUser    uuid.UUID `json:"id_user"`
	IDPlanet  uuid.UUID `json:"id_planet"`
	Type      int       `json:"RequestType"`
	Name      int       `json:"Name"`
	TimeStart int       `json:"time_start"`
	TimeNeed  int       `json:"time_need"`
	Progress  int       `json:"progress"`
}
