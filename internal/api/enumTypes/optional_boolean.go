package enumTypes

import (
	"encoding/json"
	"fmt"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/pkg/logutils"
	"strings"
)

type OptionalBoolean string

const (
	OptionalBoolean1    OptionalBoolean = "1"
	OptionalBooleanYes  OptionalBoolean = "yes"
	OptionalBooleanTrue OptionalBoolean = "true"

	OptionalBoolean0     OptionalBoolean = "0"
	OptionalBooleanNo    OptionalBoolean = "no"
	OptionalBooleanFalse OptionalBoolean = "false"

	OptionalBooleanUnknown OptionalBoolean = ""
)

func AllOptionalBooleans() containers.Set[OptionalBoolean] {
	return containers.NewSet[OptionalBoolean](
		OptionalBoolean1,
		OptionalBooleanYes,
		OptionalBooleanTrue,
		OptionalBoolean0,
		OptionalBooleanNo,
		OptionalBooleanFalse,
	)
}

func (g OptionalBoolean) String() string {
	switch g {
	case OptionalBoolean0:
		return "0"
	case OptionalBooleanNo:
		return "no"
	case OptionalBooleanFalse:
		return "false"
	case OptionalBoolean1:
		return "1"
	case OptionalBooleanYes:
		return "yes"
	case OptionalBooleanTrue:
		return "true"
	case OptionalBooleanUnknown:
		return ""
	default:
		return ""
	}
}

func TranslateOptionalBoolean(str string) (*bool, error) {
	switch strings.ToLower(str) {
	case
		string(OptionalBooleanNo), OptionalBooleanNo.String(),
		string(OptionalBooleanFalse), OptionalBooleanFalse.String(),
		string(OptionalBoolean0), OptionalBoolean0.String():
		return NewBool(false), nil
	case
		string(OptionalBoolean1), OptionalBoolean1.String(),
		string(OptionalBooleanTrue), OptionalBooleanTrue.String(),
		string(OptionalBooleanYes), OptionalBooleanYes.String():
		return NewBool(true), nil
	case string(OptionalBooleanUnknown), OptionalBooleanUnknown.String():
		return nil, nil
	default:
		return nil, fmt.Errorf("unknown value for optional boolean: \"%v\"", logutils.Escape(str))
	}
}
func ParseOptionalBoolean(str string) (OptionalBoolean, error) {
	switch strings.ToLower(str) {
	case
		string(OptionalBooleanNo), OptionalBooleanNo.String(),
		string(OptionalBooleanFalse), OptionalBooleanFalse.String(),
		string(OptionalBoolean0), OptionalBoolean0.String():
		return OptionalBooleanFalse, nil
	case
		string(OptionalBoolean1), OptionalBoolean1.String(),
		string(OptionalBooleanTrue), OptionalBooleanTrue.String(),
		string(OptionalBooleanYes), OptionalBooleanYes.String():
		return OptionalBooleanTrue, nil
	case string(OptionalBooleanUnknown), OptionalBooleanUnknown.String():
		return OptionalBooleanUnknown, nil
	default:
		return "", fmt.Errorf("unknown value for optional boolean: \"%v\"", logutils.Escape(str))
	}
}

func (g OptionalBoolean) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", string(g))), nil
}

func (g *OptionalBoolean) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	ds, err := ParseOptionalBoolean(str)
	if err != nil {
		return err
	}
	*g = ds
	return nil
}

func (g OptionalBoolean) MarshalText() ([]byte, error) {
	return []byte(g), nil
}

func (g *OptionalBoolean) UnmarshalText(b []byte) error {
	parsed, err := ParseOptionalBoolean(string(b))
	if err != nil {
		return err
	}
	*g = parsed
	return nil
}

func NewBool(b bool) *bool {
	return &b
}
