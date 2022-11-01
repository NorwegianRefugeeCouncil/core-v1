package forms

import (
	"encoding"
	"fmt"
	"strconv"
	"time"
)

// Serializer is the interface for serializing values
type Serializer interface {
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
	Serializer
	Decoder
}

type StringCodec struct{}

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

func (c *StringCodec) Decode(value string) (interface{}, error) {
	return value, nil
}

type BoolCodec struct{}

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

type IntCodec struct{}

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

type TimeCodec struct {
	format string
}

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
