package validator

import (
	"github.com/go-playground/validator/v10"
)

var isNumber validator.Func = func(fl validator.FieldLevel) bool {
	_, ok1 := fl.Field().Interface().(int)
  _, ok2 := fl.Field().Interface().(int64)
	return ok1 || ok2
}
