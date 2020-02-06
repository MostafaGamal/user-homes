package dtos

import (
	"github.com/revel/revel"
	"olarm/app/validations"
)

type Sensor struct {
	ID		string	`json:"ID"`
	Name	string	`json:"name"`
}

type HomeDTO struct {
	ID			string		`json:"ID"`
	Serial		string		`json:"serial"`
	Sensors		[]Sensor	`json:"sensors"`
	UserID		string		`json:"user_id"`
}

func (homeData *HomeDTO) Validate(v *revel.Validation) {
	validations.ValidateID(v, homeData.ID)
	validations.ValidateSerial(v, homeData.Serial)
	for _, senor := range homeData.Sensors {
		validations.ValidateID(v, senor.ID)
		validations.ValidateName(v, senor.Name)
	}
	validations.ValidateID(v, homeData.UserID)
}
