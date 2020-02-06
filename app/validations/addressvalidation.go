package validations

import "github.com/revel/revel"

func ValidateAddress(v *revel.Validation, address string) {
	v.Required(address).Message("address field is required")
}
