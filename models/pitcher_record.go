package models

import (
	"time"
)

type PitcherRecord struct {
	ID           uint      `gorm:"primary_key"`
	PlayerId     uint      `json:"player_id"`
	BelongTeamId uint      `json:"belong_team_id"`
	Year         uint      `json:"year"`
	Pitches      uint      `json:"pitches"`
	Win          uint      `json:"win"`
	Lose         uint      `json:"lose"`
	Era          float32   `json:"era"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
	DeletedAt    time.Time `json:"-"`
	Player       Player
}
