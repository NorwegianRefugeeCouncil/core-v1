package forms

// NumberInputField represents a number input field.
type NumberInputField struct {
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

// Ensure NumberInputField implements InputField
var _ InputField = (*NumberInputField)(nil)

// GetName implements FieldDefinition.GetName
func (f *NumberInputField) GetName() string {
	return f.Name
}

// GetStringValue implements FieldDefinition.GetStringValue
func (f *NumberInputField) GetStringValue() string {
	return f.Value
}

// SetStringValue implements FieldDefinition.SetStringValue
func (f *NumberInputField) SetStringValue(value string) {
	f.Value = value
}

// GetValue implements InputField.GetValue
func (f *NumberInputField) GetValue() (interface{}, error) {
	return f.getCodecOrDefault().Decode(f.Value)
}

// SetValue implements InputField.SetValue
func (f *NumberInputField) SetValue(value interface{}) error {
	codec := f.getCodecOrDefault()
	val, err := codec.Encode(value)
	if err != nil {
		return err
	}
	f.Value = val
	return nil
}

// SetErrors implements FieldDefinition.SetErrors
func (f *NumberInputField) SetErrors(errors []string) {
	f.Errors = errors
}

// HasErrors implements FieldDefinition.HasErrors
func (f *NumberInputField) HasErrors() bool {
	return len(f.Errors) > 0
}

// GetErrors implements FieldDefinition.GetErrors
func (f *NumberInputField) GetErrors() []string {
	return []string{}
}

// GetKind implements FieldDefinition.GetKind
func (f *NumberInputField) GetKind() FieldKind {
	return FieldKindNumberInput
}

func (f *NumberInputField) getCodecOrDefault() Codec {
	codec := f.Codec
	if codec == nil {
		codec = &IntCodec{}
	}
	return codec
}
