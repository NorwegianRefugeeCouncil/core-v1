package api

import "github.com/nrc-no/notcore/internal/containers"

type Gender string

const (
	GenderMale           Gender = "male"
	GenderFemale         Gender = "female"
	GenderOther          Gender = "other"
	GenderPreferNotToSay Gender = "prefers_not_to_say"
)

func AllGenders() containers.Set[Gender] {
	return containers.NewSet[Gender](
		GenderMale,
		GenderFemale,
		GenderOther,
		GenderPreferNotToSay,
	)
}

func (g Gender) String() string {
	switch g {
	case GenderMale:
		return "Male"
	case GenderFemale:
		return "Female"
	case GenderOther:
		return "Other"
	case GenderPreferNotToSay:
		return "Prefer not to say"
	default:
		return ""
	}
}
