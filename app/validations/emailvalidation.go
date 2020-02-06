package validations

import (
	"github.com/revel/revel"
	"regexp"
)

func ValidateEmail(v *revel.Validation, email string) {
	v.Required(email).Message("email field is required")
	v.Match(email,
		regexp.MustCompile("^(([a-zA-Z0-9_\\\\+]+(\\.?[a-zA-Z0-9_\\-\\\\+]+)*)|(\"[a-zA-Z0-9_\\\\+]+(\\.?[a-zA-Z0-9_\\-\\\\+]+)*\"))@(([a-zA-Z0-9]+([\\\\.-][a-zA-Z0-9]+)*(\\.[a-zA-Z]{2,4}))|([0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3})|(\\[[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}]))$")).
		Message("invalid email format")
}
