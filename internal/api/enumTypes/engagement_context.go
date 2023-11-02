package enumTypes

import (
	"encoding/json"
	"fmt"
	"github.com/nrc-no/notcore/internal/locales"

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

func AllEngagementContexts() containers.Set[EngagementContext] {
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
	t := locales.GetTranslator()
	switch g {
	case EngagementContextHouseVisit:
		return t("option_engagement_context_house_visit")
	case EngagementContextFieldActivity:
		return t("option_engagement_context_field_activity")
	case EngagementContextInOffice:
		return t("option_engagement_context_in_office")
	case EngagementContextRemoteChannels:
		return t("option_engagement_context_remote_channels")
	case EngagementContextReferred:
		return t("option_engagement_context_referred")
	case EngagementContextOther:
		return t("option_other")
	case EngagementContextUnspecified:
		return ""
	default:
		return ""
	}
}

func ParseEngagementContext(str string) (EngagementContext, error) {
	switch str {
	case string(EngagementContextHouseVisit), EngagementContextHouseVisit.String():
		return EngagementContextHouseVisit, nil
	case string(EngagementContextFieldActivity), EngagementContextFieldActivity.String():
		return EngagementContextFieldActivity, nil
	case string(EngagementContextInOffice), EngagementContextInOffice.String():
		return EngagementContextInOffice, nil
	case string(EngagementContextRemoteChannels), EngagementContextRemoteChannels.String():
		return EngagementContextRemoteChannels, nil
	case string(EngagementContextReferred), EngagementContextReferred.String():
		return EngagementContextReferred, nil
	case string(EngagementContextOther), EngagementContextOther.String():
		return EngagementContextOther, nil
	case string(EngagementContextUnspecified), EngagementContextUnspecified.String():
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
