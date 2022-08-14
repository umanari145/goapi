package models

import (
	"time"
)

type Player struct {
	ID         uint       `gorm:"primary_key"`
	PlayerName string     `json:"player_name"`
	BirthDate  time.Time  `json:"birth_date"`
	CreatedAt  time.Time  `json:"-"`
	UpdatedAt  time.Time  `json:"-"`
	DeletedAt  *time.Time `sql:"index"json:"-"`
}
