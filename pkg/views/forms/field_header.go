package forms

// HeaderField represents a header
type HeaderField struct {
	// Content is the content of the header.
	Content string
	// Level is the level of the header.
	Level int
}

// Ensure HeaderField implements Field
var _ Field = (*HeaderField)(nil)

// GetKind implements Field.GetKind
func (f *HeaderField) GetKind() FieldKind {
	return FieldKindHeader
}
