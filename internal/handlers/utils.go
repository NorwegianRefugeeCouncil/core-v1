package handlers

import (
	"github.com/nrc-no/notcore/internal/api"
	"golang.org/x/exp/slices"
)

func validateIndividualsExistInCountry(individualIds []string, existingIndividuals []*api.Individual, expectedCountryId string) []string {
	var invalidIndividualIds []string
	for _, individual := range existingIndividuals {
		if individual.CountryID != expectedCountryId || !slices.Contains(individualIds, individual.ID) {
			invalidIndividualIds = append(invalidIndividualIds, individual.ID)
		}
	}
	return invalidIndividualIds
}
