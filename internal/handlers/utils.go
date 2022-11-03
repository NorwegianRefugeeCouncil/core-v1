package handlers

import (
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/containers"
)

func validateIndividualsExistInCountry(individualIds []string, existingIndividuals []*api.Individual, expectedCountryId string) []string {
	individualIdsSet := containers.NewStringSet(individualIds...)
	invalidIndividualIds := containers.NewStringSet()
	for _, individual := range existingIndividuals {
		if individual.CountryID != expectedCountryId || !individualIdsSet.Contains(individual.ID) {
			invalidIndividualIds.Add(individual.ID)
		}
	}
	return invalidIndividualIds.Items()
}
