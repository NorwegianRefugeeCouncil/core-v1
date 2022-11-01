package forms

// SelectInputField represents a select input field.
type SelectInputField struct {
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
	// Options are the options of the field.
	Errors []string `json:"errors"`
	// Options are the options of the field.
	Options []SelectInputFieldOption `json:"options"`
	// Codec is the codec of the field.
	Codec Codec `json:"-"`
}

// Ensure SelectInputField implements InputField
var _ InputField = (*SelectInputField)(nil)

// GetName implements FieldDefinition.GetName
func (f *SelectInputField) GetName() string {
	return f.Name
}

// GetStringValue implements FieldDefinition.GetStringValue
func (f *SelectInputField) GetStringValue() string {
	return f.Value
}

// SetStringValue implements FieldDefinition.SetStringValue
func (f *SelectInputField) SetStringValue(value string) {
	f.Value = value
}

// GetValue implements InputField.GetValue
func (f *SelectInputField) GetValue() (interface{}, error) {
	return f.getCodecOrDefault().Decode(f.Value)
}

// SetValue implements InputField.SetValue
func (f *SelectInputField) SetValue(value interface{}) error {
	codec := f.getCodecOrDefault()
	val, err := codec.Encode(value)
	if err != nil {
		return err
	}
	f.Value = val
	return nil
}

// SetErrors implements FieldDefinition.SetErrors
func (f *SelectInputField) SetErrors(errors []string) {
	f.Errors = errors
}

// HasErrors implements FieldDefinition.HasErrors
func (f *SelectInputField) HasErrors() bool {
	return len(f.Errors) > 0
}

// GetErrors implements FieldDefinition.GetErrors
func (f *SelectInputField) GetErrors() []string {
	return []string{}
}

// GetKind implements FieldDefinition.GetKind
func (f *SelectInputField) GetKind() FieldKind {
	return FieldKindSelect
}

// SelectInputFieldOption represents an option of a select input field.
type SelectInputFieldOption struct {
	// Value is the value of the option.
	Value string
	// Label is the label of the option.
	Label string
}

func (f *SelectInputField) getCodecOrDefault() Codec {
	codec := f.Codec
	if codec == nil {
		codec = &StringCodec{}
	}
	return codec
}
