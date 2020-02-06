package validations

import "github.com/revel/revel"

func ValidateSerial(v *revel.Validation, serial string) {
	v.Required(serial).Message("serial field is required")
}
