package api

import (
	"encoding/json"
	"fmt"

	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/pkg/logutils"
)

type Gender string

const (
	GenderMale           Gender = "male"
	GenderFemale         Gender = "female"
	GenderOther          Gender = "other"
	GenderPreferNotToSay Gender = "prefers_not_to_say"

	GenderUnspecified Gender = ""
)

func AllGenders() containers.Set[Gender] {
	return containers.NewSet[Gender](
		GenderMale,
		GenderFemale,
		GenderOther,
		GenderPreferNotToSay,
	)
}

func (g Gender) String() string {
	switch g {
	case GenderMale:
		return "Male"
	case GenderFemale:
		return "Female"
	case GenderOther:
		return "Other"
	case GenderPreferNotToSay:
		return "Prefer not to say"
	case GenderUnspecified:
		return "Unspecified"
	default:
		return ""
	}
}

func ParseGender(str string) (Gender, error) {
	switch str {
	case string(GenderMale):
		return GenderMale, nil
	case string(GenderFemale):
		return GenderFemale, nil
	case string(GenderOther):
		return GenderOther, nil
	case string(GenderPreferNotToSay):
		return GenderPreferNotToSay, nil
	case string(GenderUnspecified):
		return GenderUnspecified, nil
	default:
		return "", fmt.Errorf("unknown gender type: %v", logutils.Escape(str))
	}
}

func (g Gender) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", string(g))), nil
}

func (g *Gender) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	gender, err := ParseGender(str)
	if err != nil {
		return err
	}
	*g = gender
	return nil
}

func (g Gender) MarshalText() ([]byte, error) {
	return []byte(g), nil
}

func (g *Gender) UnmarshalText(b []byte) error {
	parsed, err := ParseGender(string(b))
	if err != nil {
		return err
	}
	*g = parsed
	return nil
}
