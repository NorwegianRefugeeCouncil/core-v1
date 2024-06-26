package enumTypes

import (
	"encoding/json"
	"fmt"
	"github.com/nrc-no/notcore/internal/locales"

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
	t := locales.GetTranslator()
	switch g {
	case ContactMethodPhone:
		return t("option_contact_method_phone")
	case ContactMethodWhatsapp:
		return t("option_contact_method_whatsapp")
	case ContactMethodEmail:
		return t("option_contact_method_email")
	case ContactMethodVisit:
		return t("option_contact_method_visit")
	case ContactMethodOther:
		return t("option_other")
	case ContactMethodUnspecified:
		return ""
	default:
		return ""
	}
}

func ParseContactMethod(str string) (ContactMethod, error) {
	switch str {
	case string(ContactMethodPhone), ContactMethodPhone.String():
		return ContactMethodPhone, nil
	case string(ContactMethodWhatsapp), ContactMethodWhatsapp.String():
		return ContactMethodWhatsapp, nil
	case string(ContactMethodEmail), ContactMethodEmail.String():
		return ContactMethodEmail, nil
	case string(ContactMethodVisit), ContactMethodVisit.String():
		return ContactMethodVisit, nil
	case string(ContactMethodOther), ContactMethodOther.String():
		return ContactMethodOther, nil
	case string(ContactMethodUnspecified), ContactMethodUnspecified.String():
		return ContactMethodUnspecified, nil
	default:
		return "", fmt.Errorf(locales.GetTranslator()("error_unknown_preferred_contact_method", logutils.Escape(str)))
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
