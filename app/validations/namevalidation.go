package validations

import "github.com/revel/revel"

func ValidateName(v *revel.Validation, name string) {
	v.Required(name).Message("name field is required")
	v.MaxSize(name, 150).Message("name is limited by 150 chars")
}
