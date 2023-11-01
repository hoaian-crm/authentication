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
	return len(result) == 0
}

var mustFound validator.Func = func(fl validator.FieldLevel) bool {

	// idString, _ := utils.NumberToString(fl.Field())

	condition := map[string]interface{}{
		"id": fl.Field(),
	}
	table := fl.Param()
	result := map[string]interface{}{}

	config.Db.Table(table).Where(condition).Take(&result)
	return len(result) > 0
}
