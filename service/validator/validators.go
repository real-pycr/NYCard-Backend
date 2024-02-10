package validator

import (
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

func timing(fl validator.FieldLevel) bool {
	if date, ok := fl.Field().Interface().(time.Time); ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func qqmail(fl validator.FieldLevel) bool {
	email, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	pattern := `^[1-9][0-9]{4,}@qq\.com$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}

func qqnum(fl validator.FieldLevel) bool {
	qqnum, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	pattern := `^[1-9][0-9]{4,}$`
	matched, _ := regexp.MatchString(pattern, qqnum)
	return matched
}
