package forms

// CheckboxInputField represents a checkbox input field.
type CheckboxInputField struct {
	// Name is the name of the field.
	Name string `json:"name"`
	// DisplayName is the display name of the field.
	DisplayName string `json:"displayName"`
	// Required is true if the field is required.
	Required bool `json:"required"`
	// Value is the string value of the field.
	Value string `json:"value"`
	// Help is the help text of the field.
	Help string `json:"help"`
	// Errors are the errors of the field.
	Errors []string `json:"errors"`
	// Codec is the codec of the field.
	Codec Codec `json:"-"`
}

// Ensure CheckboxInputField implements InputField
var _ InputField = (*CheckboxInputField)(nil)

// GetName implements FieldDefinition.GetName
func (f *CheckboxInputField) GetName() string {
	return f.Name
}

// GetStringValue implements FieldDefinition.GetStringValue
func (f *CheckboxInputField) GetStringValue() string {
	return f.Value
}

// SetStringValue implements FieldDefinition.SetStringValue
func (f *CheckboxInputField) SetStringValue(value string) {
	f.Value = value
}

// GetValue implements InputField.GetValue
func (f *CheckboxInputField) GetValue() (interface{}, error) {
	return f.getCodecOrDefault().Decode(f.Value)
}

// SetValue implements InputField.SetValue
func (f *CheckboxInputField) SetValue(value interface{}) error {
	codec := f.getCodecOrDefault()
	val, err := codec.Encode(value)
	if err != nil {
		return err
	}
	f.Value = val
	return nil
}

// SetErrors implements FieldDefinition.SetErrors
func (f *CheckboxInputField) SetErrors(errors []string) {
	f.Errors = errors
}

// HasErrors implements FieldDefinition.HasErrors
func (f *CheckboxInputField) HasErrors() bool {
	return len(f.Errors) > 0
}

// GetErrors implements FieldDefinition.GetErrors
func (f *CheckboxInputField) GetErrors() []string {
	copied := make([]string, len(f.Errors))
	copy(copied, f.Errors)
	return copied
}

// GetKind implements FieldDefinition.GetKind
func (f *CheckboxInputField) GetKind() FieldKind {
	return FieldKindCheckboxInput
}

func (f *CheckboxInputField) getCodecOrDefault() Codec {
	codec := f.Codec
	if codec == nil {
		codec = &BoolCodec{}
	}
	return codec
}
