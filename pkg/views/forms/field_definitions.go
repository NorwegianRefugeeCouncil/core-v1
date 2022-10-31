package forms

func (f *FieldDefinition) getField() Field {
	switch {
	case f.IDField != nil:
		return f.IDField
	case f.Text != nil:
		return f.Text
	case f.Number != nil:
		return f.Number
	case f.Date != nil:
		return f.Date
	case f.Select != nil:
		return f.Select
	case f.Checkbox != nil:
		return f.Checkbox
	case f.TextArea != nil:
		return f.TextArea
	case f.Hidden != nil:
		return f.Hidden
	default:
		panic("unknown field type")
	}
}

// FieldDefinitions represents a list of field definitions.
type FieldDefinitions []*FieldDefinition

func NewFieldDefinitions(fields ...Field) *FieldDefinitions {
	ret := make(FieldDefinitions, 0)
	ret.Add(fields...)
	return &ret
}

// FindField returns the field with the given name or nil if not found.
func (f *FieldDefinitions) FindField(name string) Field {
	for _, field := range *f {
		if field.GetName() == name {
			return field
		}
	}
	return nil
}

// Each calls the given function for each field.
func (f *FieldDefinitions) Each(fn func(int, Field)) {
	for i, field := range *f {
		fn(i, field.getField())
	}
}

// Add adds a field to field definitions
func (f *FieldDefinitions) Add(fields ...Field) {
	for _, field := range fields {
		fieldDef := NewFieldDefinition(field)
		*f = append(*f, fieldDef)
	}
}
