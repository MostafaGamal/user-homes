package dtos

import (
	"github.com/revel/revel"
	"olarm/app/validations"
)

type UserDTO struct {
	Name		string		`json:"name"`
	Password	string		`json:"pass"`
	Email		string		`json:"email"`
	Address		string		`json:"address"`
	Age			int			`json:"age"`
}

func (userData *UserDTO) Validate(v *revel.Validation) {
	validations.ValidateName(v, userData.Name)
	validations.ValidatePassword(v, userData.Password)
	validations.ValidateEmail(v, userData.Email)
	validations.ValidateAddress(v, userData.Address)
	validations.ValidateAge(v, userData.Age)
}
