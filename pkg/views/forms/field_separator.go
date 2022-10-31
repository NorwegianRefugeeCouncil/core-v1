package forms

type Separator struct{}

func (s Separator) GetKind() FieldKind {
	return FieldKindSeparator
}

var _ Field = &Separator{}
