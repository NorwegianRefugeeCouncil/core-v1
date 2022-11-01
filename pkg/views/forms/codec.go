package forms

import (
	"encoding"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Encoder is the interface for serializing values
type Encoder interface {
	// Encode encodes the value to a string
	Encode(value interface{}) (string, error)
}

// Decoder is the interface for decoding values
type Decoder interface {
	// Decode decodes the value from a string
	Decode(value string) (interface{}, error)
}

// Codec represents a serializer/deserializer for a field value.
type Codec interface {
	Encoder
	Decoder
}

type codec struct {
	encoder Encoder
	decoder Decoder
}

// NewCodec returns a new codec with the given encoder and decoder
func NewCodec(encoder Encoder, decoder Decoder) Codec {
	return &codec{
		encoder: encoder,
		decoder: decoder,
	}
}

// Encode implements Encoder.Encode
func (c codec) Encode(value interface{}) (string, error) {
	return c.encoder.Encode(value)
}

// Decode implements Decoder.Decode
func (c codec) Decode(value string) (interface{}, error) {
	return c.decoder.Decode(value)
}

// StringCodec is a codec for string values
type StringCodec struct{}

// Encode implements Encoder.Encode
func (c *StringCodec) Encode(value interface{}) (string, error) {
	switch v := value.(type) {
	case encoding.TextMarshaler:
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

// BoolCodec is a codec for boolean values
type BoolCodec struct{}

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

// IntCodec is a codec for integer values
type IntCodec struct{}

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

// TimeCodec is a codec for time.Time values
type TimeCodec struct {
	format string
}

// Encode implements Encoder.Encode
func (c *TimeCodec) Encode(value interface{}) (string, error) {
	switch v := value.(type) {
	case *time.Time:
		if v == nil {
			return "", nil
		}
		return v.Format(c.format), nil
	case time.Time:
		return v.Format(c.format), nil
	}
	return "", fmt.Errorf("invalid value type")
}

// Decode implements Decoder.Decode
func (c *TimeCodec) Decode(value string) (interface{}, error) {
	if value == "" {
		return (*time.Time)(nil), nil
	}
	parsed, err := time.Parse(c.format, value)
	if err != nil {
		return (*time.Time)(nil), err
	}
	return parsed, nil
}

// StringListCodec is a codec for []string values
type StringListCodec struct{}

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
