package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func ValidatorBinding() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("is_string", isString)
		v.RegisterValidation("is_not_empty", isNotEmpty)
		v.RegisterValidation("min_length", minLength)
		v.RegisterValidation("must_unique", mustUnique)
		v.RegisterValidation("is_number", isNumber)
	}
}
