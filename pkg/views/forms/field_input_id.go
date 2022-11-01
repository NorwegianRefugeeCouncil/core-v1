package forms

// IDField represents an ID field.
type IDField struct {
	// QRCodeURL is the URL of the QR code.
	QRCodeURL string
	// Name is the name of the field.
	Name string
	// DisplayName is the display name of the field.
	DisplayName string
	// Required is true if the field is required.
	Required bool
	// Value is the string value of the field.
	Value string
	// Help is the help text of the field.
	Help string
	// Errors are the errors of the field.
	Errors []string
	// Codec is the codec of the field.
	Codec Codec
}

// Ensure IDField implements InputField
var _ InputField = (*IDField)(nil)

// GetName implements FieldDefinition.GetName
func (f *IDField) GetName() string {
	return f.Name
}

// GetStringValue implements FieldDefinition.GetStringValue
func (f *IDField) GetStringValue() string {
	return f.Value
}

// SetStringValue implements FieldDefinition.SetStringValue
func (f *IDField) SetStringValue(value string) {
	f.Value = value
}

// GetValue implements InputField.GetValue
func (f *IDField) GetValue() (interface{}, error) {
	return f.getCodecOrDefault().Decode(f.Value)
}

// SetValue implements InputField.SetValue
func (f *IDField) SetValue(value interface{}) error {
	codec := f.getCodecOrDefault()
	val, err := codec.Encode(value)
	if err != nil {
		return err
	}
	f.Value = val
	return nil
}

// SetErrors implements FieldDefinition.SetErrors
func (f *IDField) SetErrors(errors []string) {
	f.Errors = errors
}

// HasErrors implements FieldDefinition.HasErrors
func (f *IDField) HasErrors() bool {
	return len(f.Errors) > 0
}

// GetErrors implements FieldDefinition.GetErrors
func (f *IDField) GetErrors() []string {
	return []string{}
}

// GetKind implements FieldDefinition.GetKind
func (f *IDField) GetKind() FieldKind {
	return FieldKindID
}

func (f *IDField) getCodecOrDefault() Codec {
	codec := f.Codec
	if codec == nil {
		codec = &StringCodec{}
	}
	return codec
}
