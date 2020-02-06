package validations

import (
	"github.com/revel/revel"
	"regexp"
)

func ValidatePassword(v *revel.Validation, password string) {
	v.Required(password).Message("pass field is required")
	v.MaxSize(password, 128).Message("pass field is limited by 128 chars")
	v.MinSize(password, 8).Message("pass field requires 8 chars at least")
	v.Match(password, regexp.MustCompile("[a-z]+")).Message("pass field requires 1 lower char at least")
	v.Match(password, regexp.MustCompile("[A-Z]+")).Message("pass field requires 1 upper char at least")
	v.Match(password, regexp.MustCompile("[@$!%*?&_-]+")).Message("pass field requires 1 special char at least")
}
