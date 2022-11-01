package forms

import (
	"fmt"
	"time"
)

// TimeCodec is a codec for time.Time values
type TimeCodec struct {
	format string
}

func NewTimeCodec(format string) Codec {
	return &TimeCodec{format: format}
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
