package validator

import (
	"strconv"

	"github.com/go-playground/validator/v10"
)

var isString validator.Func = func(fl validator.FieldLevel) bool {
	_, ok := fl.Field().Interface().(string)
	return ok
}

var isNotEmpty validator.Func = func(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	if len(value) == 0 {
		return false
	}
	return true
}

var minLength validator.Func = func(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	minLengthValue, err := strconv.Atoi(fl.Param())

	if err != nil {
		return false
	}

	if len(value) < minLengthValue {
		return false
	}

	return true
}
