package api_test

import (
	"github.com/nrc-no/notcore/internal/api/enumTypes"
	"github.com/nrc-no/notcore/internal/handlers"
	"testing"
	"time"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/stretchr/testify/assert"
)

var foo = time.Time(time.Date(1992, 7, 31, 0, 0, 0, 0, time.UTC))

var parameters = []struct {
	column string
	value  string
	out    any
	error  bool
}{
	{constants.FileColumnIndividualID, "unique-user-id", "unique-user-id", false},
	{constants.FileColumnIndividualAddress, "123 Fake Alle, Berlin", "123 Fake Alle, Berlin", false},
	{constants.FileColumnIndividualBirthDate, "1992-07-31", &foo, false},
	{constants.FileColumnIndividualBirthDate, "31-07-1992", "", true},
	{constants.FileColumnIndividualDisplacementStatus, "refugee", enumTypes.DisplacementStatusRefugee, false},
	{constants.FileColumnIndividualHasDisability, "yEs", enumTypes.NewBool(true), false},
	{constants.FileColumnIndividualDisplacementStatus, "Refugee", enumTypes.DisplacementStatusRefugee, false},
	{constants.FileColumnIndividualDisplacementStatus, "nothing", enumTypes.DisplacementStatusRefugee, true},
	{constants.FileColumnIndividualServiceCC1, "shelter_and_settlements", enumTypes.ServiceCCShelter, false},
	{constants.FileColumnIndividualServiceCC1, "Shelter & Settlements", enumTypes.ServiceCCShelter, false},
	{constants.FileColumnIndividualServiceCC1, "nothing", enumTypes.ServiceCCShelter, true},
	{constants.FileColumnIndividualEmail1, "person@not-nrc.no", "person@not-nrc.no", false},
	{constants.FileColumnIndividualFullName, "Hugh Jazz", "Hugh Jazz", false},
	{constants.FileColumnIndividualFirstName, "Hugh", "Hugh", false},
	{constants.FileColumnIndividualMiddleName, "James", "James", false},
	{constants.FileColumnIndividualLastName, "Jazz", "Jazz", false},
	{constants.FileColumnIndividualNativeName, "جون", "جون", false},
	{constants.FileColumnIndividualMothersName, "Jane Doe", "Jane Doe", false},
	{constants.FileColumnIndividualSex, "male", enumTypes.SexMale, false},
	{constants.FileColumnIndividualIsMinor, "tRuE", enumTypes.NewBool(true), false},
	{constants.FileColumnIndividualIsMinor, "YeS", enumTypes.NewBool(true), false},
	{constants.FileColumnIndividualIsMinor, "1", enumTypes.NewBool(true), false},
	{constants.FileColumnIndividualIsMinor, "anything-else", false, true},
	{constants.FileColumnIndividualPhoneNumber1, "01234", "01234", false},
	{constants.FileColumnIndividualPreferredName, "Hughie", "Hughie", false},
	{constants.FileColumnIndividualPresentsProtectionConcerns, "tRuE", enumTypes.NewBool(true), false},
	{constants.FileColumnIndividualPresentsProtectionConcerns, "YeS", enumTypes.NewBool(true), false},
	{constants.FileColumnIndividualPresentsProtectionConcerns, "1", enumTypes.NewBool(true), false},
	{constants.FileColumnIndividualPresentsProtectionConcerns, "anything-else", "anything-else", true},
}

func TestUnmarshalIndividualsTabularData(t *testing.T) {
	// test upload limit
	var tooMuchData [][]string
	var tooManyIndividuals []*api.Individual
	var fields []string
	var uploadLimit = 20

	for _, p := range parameters {
		tooMuchData = append(tooMuchData, []string{p.column, p.value})
	}
	uploadLimitError := api.UnmarshalIndividualsTabularData(tooMuchData, &tooManyIndividuals, map[string]int{}, &uploadLimit)
	assert.Len(t, uploadLimitError, 1)
	assert.Equal(t, uploadLimitError[0].Message, "Your file contains 28 participants, which exceeds the upload limit of 20 participants at a time.")

	// test unmarshalling
	for _, param := range parameters {
		headerRow := []string{param.column}
		dataRow := []string{param.value}
		data := [][]string{headerRow, dataRow}
		colMapping, _ := api.GetColumnMapping(data, &fields)

		var individuals []*api.Individual

		fileErrors := api.UnmarshalIndividualsTabularData(data, &individuals, colMapping, &handlers.UPLOAD_LIMIT)

		if param.error {
			assert.Greater(t, len(fileErrors), 0)
		} else {
			assert.Len(t, fileErrors, 0)

			assert.Len(t, individuals, 1)
			value, err := individuals[0].GetFieldValue(param.column)
			assert.NoError(t, err)
			assert.Equal(t, param.out, value)
		}
	}
}
