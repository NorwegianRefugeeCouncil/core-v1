package validation

import (
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/utils"
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

func NormalizeIndividual(individual *api.Individual) {
	individual.FullName = utils.TrimString(individual.FullName)
	individual.PreferredName = utils.TrimString(individual.PreferredName)
	if individual.PreferredName == "" {
		individual.PreferredName = individual.FullName
	}
	individual.DisplacementStatus = utils.TrimString(individual.DisplacementStatus)
	individual.Email = utils.TrimString(utils.NormalizeEmail(individual.Email))
	individual.PhoneNumber = utils.TrimString(individual.PhoneNumber)
	individual.Address = utils.TrimString(individual.Address)
	individual.Gender = utils.TrimString(individual.Gender)
	individual.NormalizedPhoneNumber = utils.NormalizePhoneNumber(individual.PhoneNumber)
	individual.PhysicalImpairment = utils.TrimString(individual.PhysicalImpairment)
	individual.MentalImpairment = utils.TrimString(individual.MentalImpairment)
	individual.SensoryImpairment = utils.TrimString(individual.SensoryImpairment)
}
