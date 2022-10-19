package validation

import (
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/constants"
)

func ValidateIndividual(individual *api.Individual) ValidationErrors {
	ret := ValidationErrors{}
	if individual.FullName == "" {
		ret[constants.FormParamIndividualFullName] = "Full name is required"
	}
	if len(individual.CountryID) == 0 {
		ret[constants.FormParamIndividualCountry] = "Country is required"
	}
	return ret
}
