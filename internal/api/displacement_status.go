package api

import "github.com/nrc-no/notcore/internal/containers"

type DisplacementStatus string

const (
	DisplacementStatusIDP           DisplacementStatus = "idp"
	DisplacementStatusRefugee       DisplacementStatus = "refugee"
	DisplacementStatusHostCommunity DisplacementStatus = "host_community"
)

func AllDisplacementStatuses() containers.Set[DisplacementStatus] {
	return containers.NewSet[DisplacementStatus](
		DisplacementStatusIDP,
		DisplacementStatusRefugee,
		DisplacementStatusHostCommunity,
	)
}

func (g DisplacementStatus) String() string {
	switch g {
	case DisplacementStatusIDP:
		return "Internally Displaced Person"
	case DisplacementStatusRefugee:
		return "Refugee"
	case DisplacementStatusHostCommunity:
		return "Host Community"
	default:
		return ""
	}
}
