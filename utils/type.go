package utils

import (
	"reflect"
	"strconv"
	"strings"
)

func StringToNumber(value string) (int64, error) {
	result, err := strconv.Atoi(value)

	return int64(result), err
}

func SplitStringToNumbers(value string, spliter string) ([]int64, error) {

	var result []int64
	var err error

	for i, str := range strings.Split(value, spliter) {
		result[i], err = StringToNumber(str)
		if err != nil {
			return []int64{}, err
		}
	}

	return result, nil
}

func NumberToString(value int) (string, error) {
	return strconv.Itoa(value), nil
}

func ToMap(data interface{}, tagName string) map[string]interface{} {
	result := make(map[string]interface{})

	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		panic("data in ToMap function must be a struct")
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		result[t.Field(i).Name] = v.Field(i).Interface()
	}

	return result
}
