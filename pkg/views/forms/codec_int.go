package forms

import (
	"fmt"
	"strconv"
)

// IntCodec is a codec for integer values
type IntCodec struct{}

func NewIntCodec() Codec {
	return &IntCodec{}
}

// Encode implements Encoder.Encode
func (c *IntCodec) Encode(value interface{}) (string, error) {
	switch v := value.(type) {
	case *int:
		if v == nil {
			return "", nil
		}
		return strconv.Itoa(*v), nil
	case int:
		return strconv.Itoa(v), nil
	}
	return "", fmt.Errorf("invalid value type")
}

// Decode implements Decoder.Decode
func (c *IntCodec) Decode(value string) (interface{}, error) {
	if value == "" {
		return (*int)(nil), nil
	}
	ret, err := strconv.Atoi(value)
	if err != nil {
		return (*int)(nil), err
	}
	return ret, nil
}
