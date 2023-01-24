package api

import (
	"encoding/json"
	"fmt"

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
		ServiceCCNone,
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
	switch g {
	case ServiceCCNone:
		return ""
	case ServiceCCShelter:
		return "Shelter & Settlements"
	case ServiceCCWash:
		return "WASH"
	case ServiceCCProtection:
		return "Protection"
	case ServiceCCEducation:
		return "Education"
	case ServiceCCICLA:
		return "ICLA"
	case ServiceCCLFS:
		return "LFS"
	case ServiceCCCVA:
		return "CVA"
	case ServiceCCOther:
		return "Other"
	default:
		return ""
	}
}

func ParseServiceCC(str string) (ServiceCC, error) {
	switch str {
	case string(ServiceCCNone):
		return ServiceCCNone, nil
	case string(ServiceCCShelter):
		return ServiceCCShelter, nil
	case string(ServiceCCWash):
		return ServiceCCWash, nil
	case string(ServiceCCProtection):
		return ServiceCCProtection, nil
	case string(ServiceCCEducation):
		return ServiceCCEducation, nil
	case string(ServiceCCICLA):
		return ServiceCCICLA, nil
	case string(ServiceCCLFS):
		return ServiceCCLFS, nil
	case string(ServiceCCCVA):
		return ServiceCCCVA, nil
	case string(ServiceCCOther):
		return ServiceCCOther, nil
	default:
		return "", fmt.Errorf("unknown displacement status type: %v", logutils.Escape(str))
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
