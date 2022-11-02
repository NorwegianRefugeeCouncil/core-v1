package forms

import (
	"encoding"
	"fmt"
	"reflect"
)

// StringCodec is a codec for string values
type StringCodec struct{}

var _ Codec = &StringCodec{}

func NewStringCodec() Codec {
	return &StringCodec{}
}

// Encode implements Encoder.Encode
func (c *StringCodec) Encode(value interface{}) (string, error) {
	switch v := value.(type) {
	case encoding.TextMarshaler:
		if reflect.ValueOf(v).IsNil() {
			return "", nil
		}
		bytes, err := v.MarshalText()
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	case *string:
		if v == nil {
			return "", nil
		}
		return *v, nil
	case string:
		return v, nil
	case fmt.Stringer:
		return v.String(), nil
	}
	return "", fmt.Errorf("invalid value type")
}

// Decode implements Decoder.Decode
func (c *StringCodec) Decode(value string) (interface{}, error) {
	return value, nil
}
