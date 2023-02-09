package api

import (
	"encoding/json"
	"fmt"
	"github.com/nrc-no/notcore/internal/constants"

	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/pkg/logutils"
)

type ContactMethod string

const (
	ContactMethodPhone       ContactMethod = "phone"
	ContactMethodWhatsapp    ContactMethod = "whatsapp"
	ContactMethodEmail       ContactMethod = "email"
	ContactMethodVisit       ContactMethod = "visit"
	ContactMethodOther       ContactMethod = "other"
	ContactMethodUnspecified ContactMethod = ""
)

func AllContactMethods() containers.Set[ContactMethod] {
	return containers.NewSet[ContactMethod](
		ContactMethodPhone,
		ContactMethodWhatsapp,
		ContactMethodEmail,
		ContactMethodVisit,
		ContactMethodOther,
	)
}

func (g ContactMethod) String() string {
	switch g {
	case ContactMethodPhone:
		return "Phone"
	case ContactMethodWhatsapp:
		return "Whatsapp"
	case ContactMethodEmail:
		return "Email"
	case ContactMethodVisit:
		return "Visit"
	case ContactMethodOther:
		return "Other"
	case ContactMethodUnspecified:
		return ""
	default:
		return ""
	}
}

func ParseContactMethod(str string) (ContactMethod, error) {
	switch str {
	case string(ContactMethodPhone):
		return ContactMethodPhone, nil
	case string(ContactMethodWhatsapp):
		return ContactMethodWhatsapp, nil
	case string(ContactMethodEmail):
		return ContactMethodEmail, nil
	case string(ContactMethodVisit):
		return ContactMethodVisit, nil
	case string(ContactMethodOther):
		return ContactMethodOther, nil
	case string(ContactMethodUnspecified):
		return ContactMethodUnspecified, nil
	default:
		return "", fmt.Errorf("%s: invalid value \"%v\"", constants.FileColumnIndividualPreferredContactMethod, logutils.Escape(str))
	}
}

func (g ContactMethod) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", string(g))), nil
}

func (g *ContactMethod) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	ds, err := ParseContactMethod(str)
	if err != nil {
		return err
	}
	*g = ds
	return nil
}

func (g ContactMethod) MarshalText() ([]byte, error) {
	return []byte(g), nil
}

func (g *ContactMethod) UnmarshalText(b []byte) error {
	parsed, err := ParseContactMethod(string(b))
	if err != nil {
		return err
	}
	*g = parsed
	return nil
}
