package forms

import "strings"

// SelectInputField represents a select input field.
type SelectInputField struct {
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
	// Options are the options of the field.
	Errors []string
	// AllowMultiple is true if the field allows multiple values.
	AllowMultiple bool
	// Options are the options of the field.
	Options []SelectInputFieldOption
	// Codec is the codec of the field.
	Codec Codec
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

func (f *SelectInputField) IsSelected(value string) bool {
	if f.AllowMultiple {
		values := strings.Split(f.Value, ",")
		for _, v := range values {
			if v == value {
				return true
			}
		}
		return false
	}
	return f.Value == value
}

// SelectInputFieldOption represents an option of a select input field.
type SelectInputFieldOption struct {
	// Value is the value of the option.
	Value string
	// Label is the label of the option.
	Label string
	// Disabled is true if the option is enabled.
	Disabled bool
}

func (f *SelectInputField) getCodecOrDefault() Codec {
	if f.Codec != nil {
		return f.Codec
	}
	if f.AllowMultiple {
		return &StringListCodec{}
	}
	return &StringCodec{}
}
