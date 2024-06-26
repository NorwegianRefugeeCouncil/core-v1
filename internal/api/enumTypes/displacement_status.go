package enumTypes

import (
	"encoding/json"
	"fmt"
	"github.com/nrc-no/notcore/internal/locales"

	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/pkg/logutils"
)

type DisplacementStatus string

const (
	DisplacementStatusRefugee       DisplacementStatus = "refugee"
	DisplacementStatusIDP           DisplacementStatus = "idp"
	DisplacementStatusHostCommunity DisplacementStatus = "host_community"
	DisplacementStatusReturnee      DisplacementStatus = "returnee"
	DisplacementStatusAsylumSeeker  DisplacementStatus = "asylum_seeker"
	DisplacementStatusNonDisplaced  DisplacementStatus = "non_displaced"
	DisplacementStatusOther         DisplacementStatus = "other"

	DisplacementStatusUnspecified DisplacementStatus = ""
)

func AllDisplacementStatuses() containers.Set[DisplacementStatus] {
	return containers.NewSet[DisplacementStatus](
		DisplacementStatusRefugee,
		DisplacementStatusIDP,
		DisplacementStatusHostCommunity,
		DisplacementStatusReturnee,
		DisplacementStatusAsylumSeeker,
		DisplacementStatusNonDisplaced,
		DisplacementStatusOther,
	)
}

/**
List is : IDP, Refugee, Host Community, Returnee, Non-Displaced, Other
*/

func (g DisplacementStatus) String() string {
	t := locales.GetTranslator()
	switch g {
	case DisplacementStatusRefugee:
		return t("option_displacement_status_refugee")
	case DisplacementStatusIDP:
		return t("option_displacement_status_idp")
	case DisplacementStatusHostCommunity:
		return t("option_displacement_status_host_community")
	case DisplacementStatusReturnee:
		return t("option_displacement_status_returnee")
	case DisplacementStatusAsylumSeeker:
		return t("option_displacement_status_asylum_seeker")
	case DisplacementStatusNonDisplaced:
		return t("option_displacement_status_non_displaced")
	case DisplacementStatusOther:
		return t("option_other")
	case DisplacementStatusUnspecified:
		return ""
	default:
		return ""
	}
}

func ParseDisplacementStatus(str string) (DisplacementStatus, error) {
	switch str {
	case string(DisplacementStatusRefugee), DisplacementStatusRefugee.String():
		return DisplacementStatusRefugee, nil
	case string(DisplacementStatusIDP), DisplacementStatusIDP.String():
		return DisplacementStatusIDP, nil
	case string(DisplacementStatusHostCommunity), DisplacementStatusHostCommunity.String():
		return DisplacementStatusHostCommunity, nil
	case string(DisplacementStatusReturnee), DisplacementStatusReturnee.String():
		return DisplacementStatusReturnee, nil
	case string(DisplacementStatusAsylumSeeker), DisplacementStatusAsylumSeeker.String():
		return DisplacementStatusAsylumSeeker, nil
	case string(DisplacementStatusNonDisplaced), DisplacementStatusNonDisplaced.String():
		return DisplacementStatusNonDisplaced, nil
	case string(DisplacementStatusUnspecified), DisplacementStatusUnspecified.String():
		return DisplacementStatusUnspecified, nil
	case string(DisplacementStatusOther), DisplacementStatusOther.String():
		return DisplacementStatusOther, nil
	default:
		return "", fmt.Errorf(locales.GetTranslator()("error_unknown_displacement_status", logutils.Escape(str)))
	}
}

func (g DisplacementStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", string(g))), nil
}

func (g *DisplacementStatus) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	ds, err := ParseDisplacementStatus(str)
	if err != nil {
		return err
	}
	*g = ds
	return nil
}

func (g DisplacementStatus) MarshalText() ([]byte, error) {
	return []byte(g), nil
}

func (g *DisplacementStatus) UnmarshalText(b []byte) error {
	parsed, err := ParseDisplacementStatus(string(b))
	if err != nil {
		return err
	}
	*g = parsed
	return nil
}
