package enumTypes

import (
	"encoding/json"
	"fmt"
	"github.com/nrc-no/notcore/internal/locales"

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

func AllSexes() containers.Set[Sex] {
	return containers.NewSet[Sex](
		SexMale,
		SexFemale,
		SexOther,
		SexPreferNotToSay,
	)
}

func (g Sex) String() string {
	t := locales.GetTranslator()
	switch g {
	case SexMale:
		return t("option_sex_male")
	case SexFemale:
		return t("option_sex_female")
	case SexOther:
		return t("option_other")
	case SexPreferNotToSay:
		return t("option_sex_prefers_not_to_say")
	case SexUnspecified:
		return ""
	default:
		return ""
	}
}

func ParseSex(str string) (Sex, error) {
	switch str {
	case string(SexMale), SexMale.String():
		return SexMale, nil
	case string(SexFemale), SexFemale.String():
		return SexFemale, nil
	case string(SexOther), SexOther.String():
		return SexOther, nil
	case string(SexPreferNotToSay), SexPreferNotToSay.String():
		return SexPreferNotToSay, nil
	case string(SexUnspecified), SexUnspecified.String():
		return SexUnspecified, nil
	default:
		return "", fmt.Errorf(locales.GetTranslator()("error_unknown_sex", logutils.Escape(str)))
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
