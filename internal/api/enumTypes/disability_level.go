package enumTypes

import (
	"encoding/json"
	"fmt"
	"github.com/nrc-no/notcore/internal/locales"

	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/pkg/logutils"
)

type DisabilityLevel string

const (
	DisabilityLevelNone     DisabilityLevel = "none"
	DisabilityLevelMild     DisabilityLevel = "mild"
	DisabilityLevelModerate DisabilityLevel = "moderate"
	DisabilityLevelSevere   DisabilityLevel = "severe"

	DisabilityLevelUnspecified DisabilityLevel = ""
)

func AllDisabilityLevels() containers.Set[DisabilityLevel] {
	return containers.NewSet[DisabilityLevel](
		DisabilityLevelNone,
		DisabilityLevelMild,
		DisabilityLevelModerate,
		DisabilityLevelSevere,
	)
}

func (g DisabilityLevel) String() string {
	var t = locales.GetTranslator()
	switch g {
	case DisabilityLevelNone:
		return t("option_disability_none")
	case DisabilityLevelMild:
		return t("option_disability_mild")
	case DisabilityLevelModerate:
		return t("option_disability_moderate")
	case DisabilityLevelSevere:
		return t("option_disability_severe")
	case DisabilityLevelUnspecified:
		return ""
	default:
		return ""
	}
}

func ParseDisabilityLevel(str string) (DisabilityLevel, error) {
	switch str {
	case string(DisabilityLevelNone), DisabilityLevelNone.String():
		return DisabilityLevelNone, nil
	case string(DisabilityLevelMild), DisabilityLevelMild.String():
		return DisabilityLevelMild, nil
	case string(DisabilityLevelModerate), DisabilityLevelModerate.String():
		return DisabilityLevelModerate, nil
	case string(DisabilityLevelSevere), DisabilityLevelSevere.String():
		return DisabilityLevelSevere, nil
	case string(DisabilityLevelUnspecified), DisabilityLevelUnspecified.String():
		return DisabilityLevelUnspecified, nil
	default:
		return "", fmt.Errorf("unknown disability level: \"%v\"", logutils.Escape(str))
	}
}

func (g DisabilityLevel) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", string(g))), nil
}

func (g *DisabilityLevel) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	ds, err := ParseDisabilityLevel(str)
	if err != nil {
		return err
	}
	*g = ds
	return nil
}

func (g DisabilityLevel) MarshalText() ([]byte, error) {
	return []byte(g), nil
}

func (g *DisabilityLevel) UnmarshalText(b []byte) error {
	parsed, err := ParseDisabilityLevel(string(b))
	if err != nil {
		return err
	}
	*g = parsed
	return nil
}
