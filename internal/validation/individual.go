package validation

import (
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/constants"
)

func ValidateIndividual(i *api.Individual) ValidationErrors {
	ret := ValidationErrors{}
	if len(i.CountryID) == 0 {
		ret[constants.FormParamIndividualCountryID] = "Country is required"
	}
	if !i.PrefersToRemainAnonymous {
		if i.FullName == "" {
			ret[constants.FormParamIndividualFullName] = "Full name is required if individual does not prefer to remain anonymous"
		}
	}
	return ret
}
