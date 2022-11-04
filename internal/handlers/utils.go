package handlers

import (
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/containers"
)

func validateIndividualsExistInCountry(individualIds containers.StringSet, existingIndividuals []*api.Individual, expectedCountryId string) []string {
	existingIndividualIdMap := map[string]*api.Individual{}
	for _, individual := range existingIndividuals {
		existingIndividualIdMap[individual.ID] = individual
	}

	invalidIndividualIds := containers.NewStringSet()
	for _, individualId := range individualIds.Items() {
		existingIndividual, ok := existingIndividualIdMap[individualId]
		if !ok {
			invalidIndividualIds.Add(individualId)
			continue
		}
		if existingIndividual.CountryID != expectedCountryId {
			invalidIndividualIds.Add(individualId)
		}
	}

	return invalidIndividualIds.Items()
}
