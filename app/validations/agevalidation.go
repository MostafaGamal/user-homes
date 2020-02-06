package validations

import "github.com/revel/revel"

func ValidateAge(v *revel.Validation, age int) {
	v.Required(age).Message("age field is required")
}
