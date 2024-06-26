package enumTypes

import (
	"encoding/json"
	"fmt"
	"github.com/nrc-no/notcore/internal/locales"

	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/pkg/logutils"
)

type ServiceCC string

const (
	ServiceCCNone       ServiceCC = ""
	ServiceCCShelter    ServiceCC = "shelter_and_settlements"
	ServiceCCWash       ServiceCC = "wash"
	ServiceCCProtection ServiceCC = "protection"
	ServiceCCEducation  ServiceCC = "education"
	ServiceCCICLA       ServiceCC = "icla"
	ServiceCCLFS        ServiceCC = "lfs"
	ServiceCCCVA        ServiceCC = "cva"
	ServiceCCOther      ServiceCC = "other"
)

func AllServiceCCs() containers.Set[ServiceCC] {
	return containers.NewSet[ServiceCC](
		ServiceCCShelter,
		ServiceCCWash,
		ServiceCCProtection,
		ServiceCCEducation,
		ServiceCCICLA,
		ServiceCCLFS,
		ServiceCCCVA,
		ServiceCCOther,
	)
}

/**
List is : IDP, Refugee, Host Community, Returnee, Non-Displaced, Other
*/

func (g ServiceCC) String() string {
	t := locales.GetTranslator()
	switch g {
	case ServiceCCShelter:
		return t("option_service_shelter")
	case ServiceCCWash:
		return t("option_service_wash")
	case ServiceCCProtection:
		return t("option_service_protection")
	case ServiceCCEducation:
		return t("option_service_education")
	case ServiceCCICLA:
		return t("option_service_icla")
	case ServiceCCLFS:
		return t("option_service_lfs")
	case ServiceCCCVA:
		return t("option_service_cva")
	case ServiceCCOther:
		return t("option_other")
	case ServiceCCNone:
		return ""
	default:
		return ""
	}
}

func ParseServiceCC(str string) (ServiceCC, error) {
	switch str {
	case string(ServiceCCNone), ServiceCCNone.String():
		return ServiceCCNone, nil
	case string(ServiceCCShelter), ServiceCCShelter.String():
		return ServiceCCShelter, nil
	case string(ServiceCCWash), ServiceCCWash.String():
		return ServiceCCWash, nil
	case string(ServiceCCProtection), ServiceCCProtection.String():
		return ServiceCCProtection, nil
	case string(ServiceCCEducation), ServiceCCEducation.String():
		return ServiceCCEducation, nil
	case string(ServiceCCICLA), ServiceCCICLA.String():
		return ServiceCCICLA, nil
	case string(ServiceCCLFS), ServiceCCLFS.String():
		return ServiceCCLFS, nil
	case string(ServiceCCCVA), ServiceCCCVA.String():
		return ServiceCCCVA, nil
	case string(ServiceCCOther), ServiceCCOther.String():
		return ServiceCCOther, nil
	default:
		return "", fmt.Errorf(locales.GetTranslator()("error_unknown_service_type", logutils.Escape(str)))
	}
}

func (g ServiceCC) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", string(g))), nil
}

func (g *ServiceCC) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	ds, err := ParseServiceCC(str)
	if err != nil {
		return err
	}
	*g = ds
	return nil
}

func (g ServiceCC) MarshalText() ([]byte, error) {
	return []byte(g), nil
}

func (g *ServiceCC) UnmarshalText(b []byte) error {
	parsed, err := ParseServiceCC(string(b))
	if err != nil {
		return err
	}
	*g = parsed
	return nil
}
