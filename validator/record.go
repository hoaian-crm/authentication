package validator

import (
	"main/config"
	// "reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var mustUnique validator.Func = func(fl validator.FieldLevel) bool {
	condition := map[string]interface{}{
		strings.ToLower(fl.FieldName()): fl.Field(),
	}
	table := fl.Param()
	result := map[string]interface{}{}

	config.Db.Table(table).Where(condition).Take(&result)
	if len(result) == 0 {
		return true
	}

	return false
}
