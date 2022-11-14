package api

import (
	"encoding/json"
	"fmt"

	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/pkg/logutils"
)

type Sex string

const (
	SexMale           Sex = "male"
	SexFemale         Sex = "female"
	SexOther          Sex = "other"
	SexPreferNotToSay Sex = "prefers_not_to_say"

	SexUnspecified Sex = ""
)

func AllSexs() containers.Set[Sex] {
	return containers.NewSet[Sex](
		SexMale,
		SexFemale,
		SexOther,
		SexPreferNotToSay,
	)
}

func (g Sex) String() string {
	switch g {
	case SexMale:
		return "Male"
	case SexFemale:
		return "Female"
	case SexOther:
		return "Other"
	case SexPreferNotToSay:
		return "Prefer not to say"
	case SexUnspecified:
		return "Unspecified"
	default:
		return ""
	}
}

func ParseSex(str string) (Sex, error) {
	switch str {
	case string(SexMale):
		return SexMale, nil
	case string(SexFemale):
		return SexFemale, nil
	case string(SexOther):
		return SexOther, nil
	case string(SexPreferNotToSay):
		return SexPreferNotToSay, nil
	case string(SexUnspecified):
		return SexUnspecified, nil
	default:
		return "", fmt.Errorf("unknown sex type: %v", logutils.Escape(str))
	}
}

func (g Sex) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", string(g))), nil
}

func (g *Sex) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	sex, err := ParseSex(str)
	if err != nil {
		return err
	}
	*g = sex
	return nil
}

func (g Sex) MarshalText() ([]byte, error) {
	return []byte(g), nil
}

func (g *Sex) UnmarshalText(b []byte) error {
	parsed, err := ParseSex(string(b))
	if err != nil {
		return err
	}
	*g = parsed
	return nil
}
