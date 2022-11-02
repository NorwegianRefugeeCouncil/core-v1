package forms

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
