package models

import (
	"time"
)

type BatterRecord struct {
	ID             uint      `gorm:"primary_key"`
	PlayerId       uint      `json:"user_id"`
	BelongTeamId   uint      `json:"belong_team_id"`
	Year           uint      `json:"year"`
	AtBats         uint      `json:"at_bats"`
	Hit            uint      `json:"hit"`
	BattingAverage uint      `json:"batting_average"`
	HomeRun        uint      `json:"home_run"`
	Rbi            uint      `json:"rbi"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
	DeletedAt      time.Time `json:"-"`
	Player         Player
}
