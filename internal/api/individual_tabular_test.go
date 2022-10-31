package api_test

import (
	"testing"
	"time"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/stretchr/testify/assert"
)

var foo = time.Time(time.Date(1992, 7, 31, 0, 0, 0, 0, time.UTC))

var paramaters = []struct {
	column string
	value  string
	out    any
	error  bool
}{
	{constants.FileColumnIndividualID, "unique-user-id", "unique-user-id", false},
	{constants.FileColumnIndividualAddress, "123 Fake Alle, Berlin", "123 Fake Alle, Berlin", false},
	{constants.FileColumnIndividualBirthDate, "1992-07-31", &foo, false},
	{constants.FileColumnIndividualBirthDate, "31-07-1992", "", true},
	{constants.FileColumnIndividualDisplacementStatus, "refugee", api.DisplacementStatusRefugee, false},
	{constants.FileColumnIndividualEmail, "person@not-nrc.no", "person@not-nrc.no", false},
	{constants.FileColumnIndividualFullName, "Hugh Jazz", "Hugh Jazz", false},
	{constants.FileColumnIndividualGender, "Other", api.Gender("Other"), false},
	{constants.FileColumnIndividualIsMinor, "tRuE", true, false},
	{constants.FileColumnIndividualIsMinor, "YeS", true, false},
	{constants.FileColumnIndividualIsMinor, "1", true, false},
	{constants.FileColumnIndividualIsMinor, "anything-else", false, false},
	{constants.FileColumnIndividualPhoneNumber, "01234", "01234", false},
	{constants.FileColumnIndividualPreferredName, "Hughie", "Hughie", false},
	{constants.FileColumnIndividualPresentsProtectionConcerns, "tRuE", true, false},
	{constants.FileColumnIndividualPresentsProtectionConcerns, "YeS", true, false},
	{constants.FileColumnIndividualPresentsProtectionConcerns, "1", true, false},
	{constants.FileColumnIndividualPresentsProtectionConcerns, "anything-else", false, false},
}

func TestUnmarshalIndividualsTabularData(t *testing.T) {
	for _, param := range paramaters {
		headerRow := []string{param.column}
		dataRow := []string{param.value}
		data := [][]string{headerRow, dataRow}

		var individuals []*api.Individual
		var fields []string

		err := api.UnmarshalIndividualsTabularData(data, &individuals, &fields)

		if param.error {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)

			assert.Len(t, individuals, 1)
			value, err := individuals[0].GetFieldValue(param.column)
			assert.NoError(t, err)
			assert.Equal(t, param.out, value)

			assert.Len(t, fields, 1)
			assert.Equal(t, constants.IndividualFileToDBMap[param.column], fields[0])
		}
	}
}
