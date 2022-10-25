package validation

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/nrc-no/notcore/internal/containers"
)

// ValidationErrors is a map of field names to error messages.
type ValidationErrors map[string]string

type omitValueType struct{}

var omitValue = omitValueType{}

type nilValueType struct{}

var nilValue = nilValueType{}

type Error struct {
	Type     ErrorType
	Field    string
	BadValue interface{}
	Detail   string
}

var _ error = &Error{}

func (v *Error) Error() string {
	if v == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%s: %s", v.Field, v.ErrorBody())
}

func (v *Error) ErrorBody() string {
	var s string
	switch {
	case v.Type == ErrorTypeRequired:
		s = v.Type.String()
	case v.Type == ErrorTypeForbidden:
		s = v.Type.String()
	case v.Type == ErrorTypeTooLong:
		s = v.Type.String()
	case v.Type == ErrorTypeTooShort:
		s = v.Type.String()
	case v.Type == ErrorTypeInternal:
		s = v.Type.String()
	case v.BadValue == omitValue:
		s = v.Type.String()
	default:
		value := v.BadValue
		valueType := reflect.TypeOf(value)
		if value == nil || valueType == nil {
			value = nilValue
		} else if valueType.Kind() == reflect.Pointer {
			if reflectValue := reflect.ValueOf(value); reflectValue.IsNil() {
				value = nilValue
			} else {
				value = reflectValue.Elem().Interface()
			}
		}
		switch t := value.(type) {
		case int64, int32, float64, float32, bool:
			s = fmt.Sprintf("%s: %v", v.Type, t)
		case string:
			s = fmt.Sprintf("%s: %q", v.Type, t)
		case nilValueType:
			s = fmt.Sprintf("%s: null", v.Type)
		case fmt.Stringer:
			s = fmt.Sprintf("%s: %q", v.Type, t.String())
		default:
			s = fmt.Sprintf("%s: %#v", v.Type, value)
		}
	}
	if len(v.Detail) != 0 {
		s += fmt.Sprintf(": %s", v.Detail)
	}
	return s
}

type ErrorType string

const (
	ErrorTypeNotFound         ErrorType = "FieldValueNotFound"
	ErrorTypeRequired         ErrorType = "FieldValueRequired"
	ErrorTypeDuplicate        ErrorType = "FieldValueDuplicate"
	ErrorTypeInvalid          ErrorType = "FieldValueInvalid"
	ErrorTypeNotSupported     ErrorType = "FieldValueNotSupported"
	ErrorTypeForbidden        ErrorType = "FieldValueForbidden"
	ErrorTypeTooLong          ErrorType = "FieldValueTooLong"
	ErrorTypeTooShort         ErrorType = "FieldValueTooShort"
	ErrorTypeTooMany          ErrorType = "FieldValueTooMany"
	ErrorTypeTooFew           ErrorType = "FieldValueTooFew"
	ErrorTypeInternal         ErrorType = "InternalError"
	ErrorTypeValueTypeInvalid ErrorType = "FieldValueTypeInvalid"
)

func (t ErrorType) String() string {
	switch t {
	case ErrorTypeNotFound:
		return "Not Found"
	case ErrorTypeRequired:
		return "Required"
	case ErrorTypeDuplicate:
		return "Duplicate value"
	case ErrorTypeInvalid:
		return "Invalid value"
	case ErrorTypeNotSupported:
		return "Unsupported value"
	case ErrorTypeForbidden:
		return "Forbidden"
	case ErrorTypeTooLong:
		return "Too long"
	case ErrorTypeTooShort:
		return "Too short"
	case ErrorTypeTooMany:
		return "Too many"
	case ErrorTypeTooFew:
		return "Too few"
	case ErrorTypeInternal:
		return "Internal error"
	case ErrorTypeValueTypeInvalid:
		return "Invalid value type"
	default:
		panic(fmt.Sprintf("unknown error type %q", string(t)))
	}
}

func TypeInvalid(field *Path, value interface{}, detail string) *Error {
	return &Error{ErrorTypeValueTypeInvalid, field.String(), value, detail}
}

func NotFound(field *Path, value interface{}) *Error {
	return &Error{ErrorTypeNotFound, field.String(), value, ""}
}

func Required(field *Path, detail string) *Error {
	return &Error{ErrorTypeRequired, field.String(), "", detail}
}

func Duplicate(field *Path, value interface{}, detail string) *Error {
	return &Error{ErrorTypeDuplicate, field.String(), value, detail}
}

func Invalid(field *Path, value interface{}, detail string) *Error {
	return &Error{ErrorTypeInvalid, field.String(), value, detail}
}

func NotSupported(field *Path, value interface{}, validValues []string) *Error {
	detail := ""
	if len(validValues) > 0 {
		quotedValues := make([]string, len(validValues))
		for i, v := range validValues {
			quotedValues[i] = strconv.Quote(v)
		}
		detail = fmt.Sprintf("supported values: %s", strings.Join(quotedValues, ", "))
	}
	return &Error{ErrorTypeNotSupported, field.String(), value, detail}
}

func Forbidden(field *Path, detail string) *Error {
	return &Error{ErrorTypeForbidden, field.String(), "", detail}
}

func TooLong(field *Path, value interface{}, maxLength int) *Error {
	return &Error{ErrorTypeTooLong, field.String(), value, fmt.Sprintf("must have at most %d bytes", maxLength)}
}

func TooShort(field *Path, value interface{}, minLength int) *Error {
	return &Error{ErrorTypeTooShort, field.String(), value, fmt.Sprintf("must have at least %d bytes", minLength)}
}

func TooLongMaxLength(field *Path, value interface{}, maxLength int) *Error {
	var msg string
	if maxLength >= 0 {
		msg = fmt.Sprintf("value may not be longer than %d", maxLength)
	} else {
		msg = "value is too long"
	}
	return &Error{ErrorTypeTooLong, field.String(), value, msg}
}

func TooShortMinLength(field *Path, value interface{}, minLength int) *Error {
	var msg string
	if minLength >= 0 {
		msg = fmt.Sprintf("value may not be shorter than %d", minLength)
	} else {
		msg = "value is too short"
	}
	return &Error{ErrorTypeTooShort, field.String(), value, msg}
}

func TooMany(field *Path, actualQuantity, maxQuantity int) *Error {
	var msg = ""
	if maxQuantity >= 0 {
		msg = fmt.Sprintf("must have at most %d item", maxQuantity)
		if maxQuantity > 1 {
			msg += "s"
		}
	} else {
		msg = "too many items"
	}
	var actual interface{}
	if actualQuantity >= 0 {
		actual = actualQuantity
	} else {
		actual = omitValue
	}

	return &Error{ErrorTypeTooMany, field.String(), actual, msg}
}

func TooFew(field *Path, actualQuantity, minQuantity int) *Error {
	var msg = ""
	if minQuantity >= 0 {
		msg = fmt.Sprintf("must have at least %d item", minQuantity)
		if minQuantity > 1 {
			msg += "s"
		}
	} else {
		msg = "too few items"
	}
	var actual interface{}
	if actualQuantity >= 0 {
		actual = actualQuantity
	} else {
		actual = omitValue
	}
	return &Error{ErrorTypeTooFew, field.String(), actual, msg}
}

func InternalError(field *Path, err error) *Error {
	msg := "<nil>"
	if err != nil {
		msg = err.Error()
	}
	return &Error{ErrorTypeInternal, field.String(), "", msg}
}

type ErrorList []*Error

type Aggregate interface {
	error
	Errors() []error
	Is(error) bool
}

func NewAggregate(errList []error) Aggregate {
	if len(errList) == 0 {
		return nil
	}
	var errs []error
	for _, e := range errList {
		if e != nil {
			errs = append(errs, e)
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return aggregate(errs)
}

type aggregate []error

func (agg aggregate) Is(target error) bool {
	return agg.visit(func(err error) bool {
		return errors.Is(err, target)
	})
}

func (agg aggregate) Error() string {
	if len(agg) == 0 {
		return ""
	}
	if len(agg) == 1 {
		return agg[0].Error()
	}
	seenErrors := containers.NewStringSet()
	result := ""
	agg.visit(func(err error) bool {
		msg := err.Error()
		if seenErrors.Contains(msg) {
			return false
		}
		seenErrors.Add(msg)
		if seenErrors.Len() > 1 {
			result += ", "
		}
		result += msg
		return false
	})
	if seenErrors.Len() == 1 {
		return result
	}
	return "[" + result + "]"
}

func (agg aggregate) visit(f func(err error) bool) bool {
	for _, err := range agg {
		switch err := err.(type) {
		case aggregate:
			if match := err.visit(f); match {
				return true
			}
		case Aggregate:
			for _, nestedErr := range err.Errors() {
				if match := f(nestedErr); match {
					return true
				}
			}
		default:
			if match := f(err); match {
				return true
			}
		}
	}
	return false
}

func (agg aggregate) Errors() []error {
	return agg
}
