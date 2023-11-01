package validator

import (
	"main/config"
	"main/utils"

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
	return len(result) == 0
}

var mustFound validator.Func = func(fl validator.FieldLevel) bool {

	idString, _ := utils.NumberToString(int(fl.Field().Int()))

	condition := map[string]interface{}{
		"id": idString,
	}
	table := fl.Param()
	result := map[string]interface{}{}

	config.Db.Table(table).Where(condition).Take(&result)
	return len(result) > 0
}
