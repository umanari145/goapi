package models

type Team struct {
	ID       uint   `gorm:"primary_key"`
	TeamName string `json:"team_name"`
}
