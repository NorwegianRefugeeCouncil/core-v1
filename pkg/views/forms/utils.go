package forms

import "reflect"

func ensurePtr(value interface{}) interface{} {
	isPtr := reflect.ValueOf(value).Kind() == reflect.Ptr
	if !isPtr {
		val := reflect.ValueOf(value)
		ptr := reflect.New(val.Type())
		ptr.Elem().Set(val)
		return ptr.Interface()
	}
	return value
}
