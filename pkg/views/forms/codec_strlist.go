package forms

import (
	"fmt"
	"strings"
)

// StringListCodec is a codec for []string values
type StringListCodec struct{}

var _ Codec = &StringListCodec{}

func NewStringListCodec() Codec {
	return &StringListCodec{}
}

// Encode implements Encoder.Encode
func (c *StringListCodec) Encode(value interface{}) (string, error) {
	switch v := value.(type) {
	case []string:
		return strings.Join(v, ","), nil
	}
	return "", fmt.Errorf("invalid value type")
}

// Decode implements Decoder.Decode
func (c *StringListCodec) Decode(value string) (interface{}, error) {
	if value == "" {
		return []string{}, nil
	}
	return strings.Split(value, ","), nil
}
