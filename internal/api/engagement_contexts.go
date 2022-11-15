package api

import (
	"encoding/json"
	"fmt"

	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/pkg/logutils"
)

type EngagementContext string

const (
	EngagementContextHouseVisit     EngagementContext = "houseVisit"
	EngagementContextFieldActivity  EngagementContext = "fieldActivity"
	EngagementContextInOffice       EngagementContext = "inOffice"
	EngagementContextRemoteChannels EngagementContext = "remoteChannels"
	EngagementContextReferred       EngagementContext = "referred"
	EngagementContextOther          EngagementContext = "other"

	EngagementContextUnspecified EngagementContext = ""
)

func AllEngagementContextes() containers.Set[EngagementContext] {
	return containers.NewSet[EngagementContext](
		EngagementContextHouseVisit,
		EngagementContextFieldActivity,
		EngagementContextInOffice,
		EngagementContextRemoteChannels,
		EngagementContextReferred,
		EngagementContextOther,
	)
}

/**
List is : House Visit, Field Activity, In-Office, Remote Channels, Referred, Other
*/

func (g EngagementContext) String() string {
	switch g {
	case EngagementContextHouseVisit:
		return "House Visit"
	case EngagementContextFieldActivity:
		return "Field Activity"
	case EngagementContextInOffice:
		return "In-Office"
	case EngagementContextRemoteChannels:
		return "Remote Channels"
	case EngagementContextReferred:
		return "Referred"
	case EngagementContextOther:
		return "Other"
	case EngagementContextUnspecified:
		return "Unspecified"
	default:
		return ""
	}
}

func ParseEngagementContext(str string) (EngagementContext, error) {
	switch str {
	case string(EngagementContextHouseVisit):
		return EngagementContextHouseVisit, nil
	case string(EngagementContextFieldActivity):
		return EngagementContextFieldActivity, nil
	case string(EngagementContextInOffice):
		return EngagementContextInOffice, nil
	case string(EngagementContextRemoteChannels):
		return EngagementContextRemoteChannels, nil
	case string(EngagementContextReferred):
		return EngagementContextReferred, nil
	case string(EngagementContextOther):
		return EngagementContextOther, nil
	case string(EngagementContextUnspecified):
		return EngagementContextUnspecified, nil
	default:
		return "", fmt.Errorf("unknown engagement context type: %v", logutils.Escape(str))
	}
}

func (g EngagementContext) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", string(g))), nil
}

func (g *EngagementContext) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	ds, err := ParseEngagementContext(str)
	if err != nil {
		return err
	}
	*g = ds
	return nil
}

func (g EngagementContext) MarshalText() ([]byte, error) {
	return []byte(g), nil
}

func (g *EngagementContext) UnmarshalText(b []byte) error {
	parsed, err := ParseEngagementContext(string(b))
	if err != nil {
		return err
	}
	*g = parsed
	return nil
}
