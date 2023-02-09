package enumTypes

import (
	"encoding/json"
	"fmt"

	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/pkg/logutils"
)

type IdentificationType string

const (
	IdentificationTypePassport IdentificationType = "passport"
	IdentificationTypeUNHCR    IdentificationType = "unhcr_id"
	IdentificationTypeNational IdentificationType = "national_id"
	IdentificationTypeOther    IdentificationType = "other"

	IdentificationTypeUnspecified IdentificationType = ""
)

func AllIdentificationTypes() containers.Set[IdentificationType] {
	return containers.NewSet[IdentificationType](
		IdentificationTypePassport,
		IdentificationTypeUNHCR,
		IdentificationTypeNational,
		IdentificationTypeOther,
	)
}

/**
List is : Passport, UNHCR ID, National ID, Other
*/

func (g IdentificationType) String() string {
	switch g {
	case IdentificationTypePassport:
		return "Passport"
	case IdentificationTypeUNHCR:
		return "UNHCR ID"
	case IdentificationTypeNational:
		return "National ID"
	case IdentificationTypeOther:
		return "Other"
	case IdentificationTypeUnspecified:
		return ""
	default:
		return ""
	}
}

func ParseIdentificationType(str string) (IdentificationType, error) {
	switch str {
	case string(IdentificationTypePassport):
		return IdentificationTypePassport, nil
	case IdentificationTypePassport.String():
		return IdentificationTypePassport, nil
	case string(IdentificationTypeUNHCR):
		return IdentificationTypeUNHCR, nil
	case IdentificationTypeUNHCR.String():
		return IdentificationTypeUNHCR, nil
	case string(IdentificationTypeNational):
		return IdentificationTypeNational, nil
	case IdentificationTypeNational.String():
		return IdentificationTypeNational, nil
	case string(IdentificationTypeUnspecified):
		return IdentificationTypeUnspecified, nil
	case IdentificationTypeUnspecified.String():
		return IdentificationTypeUnspecified, nil
	case string(IdentificationTypeOther):
		return IdentificationTypeOther, nil
	case IdentificationTypeOther.String():
		return IdentificationTypeOther, nil
	default:
		return "", fmt.Errorf("identification type: invalid type: \"%v\"", logutils.Escape(str))
	}
}

func (g IdentificationType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", string(g))), nil
}

func (g *IdentificationType) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	ds, err := ParseIdentificationType(str)
	if err != nil {
		return err
	}
	*g = ds
	return nil
}

func (g IdentificationType) MarshalText() ([]byte, error) {
	return []byte(g), nil
}

func (g *IdentificationType) UnmarshalText(b []byte) error {
	parsed, err := ParseIdentificationType(string(b))
	if err != nil {
		return err
	}
	*g = parsed
	return nil
}
