package forms

// OptionalBooleanInputField represents a checkbox input field.
type OptionalBooleanInputField struct {
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

// Ensure OptionalBooleanInputField implements InputField
var _ InputField = (*OptionalBooleanInputField)(nil)

// GetName implements FieldDefinition.GetName
func (f *OptionalBooleanInputField) GetName() string {
	return f.Name
}

// GetStringValue implements FieldDefinition.GetStringValue
func (f *OptionalBooleanInputField) GetStringValue() string {
	return f.Value
}

// SetStringValue implements FieldDefinition.SetStringValue
func (f *OptionalBooleanInputField) SetStringValue(value string) {
	f.Value = value
}

// GetValue implements InputField.GetValue
func (f *OptionalBooleanInputField) GetValue() (interface{}, error) {
	return f.getCodecOrDefault().Decode(f.Value)
}

// SetValue implements InputField.SetValue
func (f *OptionalBooleanInputField) SetValue(value interface{}) error {
	codec := f.getCodecOrDefault()
	val, err := codec.Encode(value)
	if err != nil {
		return err
	}
	f.Value = val
	return nil
}

// SetErrors implements FieldDefinition.SetErrors
func (f *OptionalBooleanInputField) SetErrors(errors []string) {
	f.Errors = errors
}

// HasErrors implements FieldDefinition.HasErrors
func (f *OptionalBooleanInputField) HasErrors() bool {
	return len(f.Errors) > 0
}

// GetErrors implements FieldDefinition.GetErrors
func (f *OptionalBooleanInputField) GetErrors() []string {
	copied := make([]string, len(f.Errors))
	copy(copied, f.Errors)
	return copied
}

func (f *OptionalBooleanInputField) IsSelected(value string) bool {
	return f.Value == value
}

// GetKind implements FieldDefinition.GetKind
func (f *OptionalBooleanInputField) GetKind() FieldKind {
	return FieldKindOptionalBooleanInput
}

func (f *OptionalBooleanInputField) getCodecOrDefault() Codec {
	codec := f.Codec
	if codec == nil {
		codec = &BoolCodec{}
	}
	return codec
}
