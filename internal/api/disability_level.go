package api

import (
	"fmt"
	"strings"
)

type DisabilityLevel uint8

const (
	DisabilityLevelNone DisabilityLevel = iota
	DisabilityLevelMild
	DisabilityLevelModerate
	DisabilityLevelSevere
)

func (d DisabilityLevel) String() string {
	switch d {
	case DisabilityLevelNone:
		return "None"
	case DisabilityLevelMild:
		return "Mild"
	case DisabilityLevelModerate:
		return "Moderate"
	case DisabilityLevelSevere:
		return "Severe"
	default:
		return "Unknown"
	}
}

func ParseDisabilityLevel(s string) (DisabilityLevel, error) {
	switch strings.ToLower(s) {
	case "0":
		return DisabilityLevelNone, nil
	case "1":
		return DisabilityLevelMild, nil
	case "2":
		return DisabilityLevelModerate, nil
	case "3":
		return DisabilityLevelSevere, nil
	default:
		return DisabilityLevelNone, fmt.Errorf("unknown disability level %s", s)
	}
}
