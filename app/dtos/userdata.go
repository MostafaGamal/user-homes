package dtos

import (
	"encoding/json"
	"olarm/app/infrastructure"
)

type UserDataDTO struct {
	User	UserDTO		`json:"userData"`
	Homes	[]HomeDTO	`json:"homes"`
}

func ToHomeDTO(home *infrastructure.Home) *HomeDTO {
	var sensors []Sensor
	err := json.Unmarshal([]byte(home.Sensors), &sensors)
	if err != nil {
		return &HomeDTO{}
	}
	return &HomeDTO{
		ID:      home.ID,
		Serial:  home.Serial,
		Sensors: sensors,
		UserID:  home.UserID,
	}
}
