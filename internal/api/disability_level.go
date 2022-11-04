package api

import (
	"encoding/json"
	"fmt"

	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/utils"
)

type DisabilityLevel string

const (
	DisabilityLevelNone     DisabilityLevel = "none"
	DisabilityLevelMild     DisabilityLevel = "mild"
	DisabilityLevelModerate DisabilityLevel = "moderate"
	DisabilityLevelSevere   DisabilityLevel = "severe"
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
	switch g {
	case DisabilityLevelNone:
		return "None"
	case DisabilityLevelMild:
		return "Mild"
	case DisabilityLevelModerate:
		return "Moderate"
	case DisabilityLevelSevere:
		return "Severe"
	default:
		return ""
	}
}

func ParseDisabilityLevel(str string) (DisabilityLevel, error) {
	switch str {
	case string(DisabilityLevelNone):
		return DisabilityLevelNone, nil
	case string(DisabilityLevelMild):
		return DisabilityLevelMild, nil
	case string(DisabilityLevelModerate):
		return DisabilityLevelModerate, nil
	case string(DisabilityLevelSevere):
		return DisabilityLevelSevere, nil
	default:
		return "", fmt.Errorf("unknown disability level: %v", utils.LogEscape(str))
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
