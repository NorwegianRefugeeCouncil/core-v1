package forms

import "reflect"

func ensurePtr(value interface{}) interface{} {
	if reflect.ValueOf(value).Kind() != reflect.Ptr {
		return &value
	}
	return value
}
