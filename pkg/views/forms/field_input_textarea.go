package forms

// TextAreaInputField represents a text area input field.
type TextAreaInputField struct {
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

// Ensure TextAreaInputField implements InputField
var _ InputField = (*TextAreaInputField)(nil)

// GetName implements FieldDefinition.GetName
func (f *TextAreaInputField) GetName() string {
	return f.Name
}

// GetStringValue implements FieldDefinition.GetStringValue
func (f *TextAreaInputField) GetStringValue() string {
	return f.Value
}

// SetStringValue implements FieldDefinition.SetStringValue
func (f *TextAreaInputField) SetStringValue(value string) {
	f.Value = value
}

// GetValue implements InputField.GetValue
func (f *TextAreaInputField) GetValue() (interface{}, error) {
	return f.getCodecOrDefault().Decode(f.Value)
}

// SetValue implements InputField.SetValue
func (f *TextAreaInputField) SetValue(value interface{}) error {
	codec := f.getCodecOrDefault()
	val, err := codec.Encode(value)
	if err != nil {
		return err
	}
	f.Value = val
	return nil
}

// SetErrors implements FieldDefinition.SetErrors
func (f *TextAreaInputField) SetErrors(errors []string) {
	f.Errors = errors
}

// HasErrors implements FieldDefinition.HasErrors
func (f *TextAreaInputField) HasErrors() bool {
	return len(f.Errors) > 0
}

// GetErrors implements FieldDefinition.GetErrors
func (f *TextAreaInputField) GetErrors() []string {
	return []string{}
}

// GetKind implements FieldDefinition.GetKind
func (f *TextAreaInputField) GetKind() FieldKind {
	return FieldKindTextarea
}

func (f *TextAreaInputField) getCodecOrDefault() Codec {
	codec := f.Codec
	if codec == nil {
		codec = &StringCodec{}
	}
	return codec
}
