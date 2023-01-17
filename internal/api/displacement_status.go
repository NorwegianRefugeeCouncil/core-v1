package api

import (
	"encoding/json"
	"fmt"

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
	switch g {
	case DisplacementStatusRefugee:
		return "Refugee"
	case DisplacementStatusIDP:
		return "IDP"
	case DisplacementStatusHostCommunity:
		return "Host Community"
	case DisplacementStatusReturnee:
		return "Returnee"
	case DisplacementStatusAsylumSeeker:
		return "Asylum Seeker"
	case DisplacementStatusNonDisplaced:
		return "Non-Displaced"
	case DisplacementStatusOther:
		return "Other"
	case DisplacementStatusUnspecified:
		return "Unspecified"
	default:
		return ""
	}
}

func ParseDisplacementStatus(str string) (DisplacementStatus, error) {
	switch str {
	case string(DisplacementStatusRefugee):
		return DisplacementStatusRefugee, nil
	case string(DisplacementStatusIDP):
		return DisplacementStatusIDP, nil
	case string(DisplacementStatusHostCommunity):
		return DisplacementStatusHostCommunity, nil
	case string(DisplacementStatusReturnee):
		return DisplacementStatusReturnee, nil
	case string(DisplacementStatusAsylumSeeker):
		return DisplacementStatusAsylumSeeker, nil
	case string(DisplacementStatusNonDisplaced):
		return DisplacementStatusNonDisplaced, nil
	case string(DisplacementStatusUnspecified):
		return DisplacementStatusUnspecified, nil
	case string(DisplacementStatusOther):
		return DisplacementStatusOther, nil
	default:
		return "", fmt.Errorf("unknown displacement status type: %v", logutils.Escape(str))
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
