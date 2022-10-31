package forms

type Spacer struct{}

func (s Spacer) GetKind() FieldKind {
	return FieldKindSpacer
}

var _ Field = &Spacer{}
