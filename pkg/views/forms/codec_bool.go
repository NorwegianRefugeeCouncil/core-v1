package forms

import (
	"fmt"
	"strconv"
)

// BoolCodec is a codec for boolean values
type BoolCodec struct{}

var _ Codec = &BoolCodec{}

func NewBoolCodec() Codec {
	return &BoolCodec{}
}

// Encode implements Encoder.Encode
func (c *BoolCodec) Encode(value interface{}) (string, error) {
	switch v := value.(type) {
	case *bool:
		if v == nil {
			return "", nil
		}
		return strconv.FormatBool(*v), nil
	case bool:
		return strconv.FormatBool(v), nil
	}
	return "", fmt.Errorf("invalid value type")
}

// Decode implements Decoder.Decode
func (c *BoolCodec) Decode(value string) (interface{}, error) {
	if value == "" {
		return (*bool)(nil), nil
	}
	ret, err := strconv.ParseBool(value)
	if err != nil {
		return (*bool)(nil), err
	}
	return ret, nil
}
