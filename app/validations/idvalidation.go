package validations

import "github.com/revel/revel"

func ValidateID(v *revel.Validation, id string) {
	v.Required(id).Message("ID field is required")
}
