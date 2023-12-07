package api

import (
	"errors"
	"github.com/nrc-no/notcore/internal/api/enumTypes"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/locales"
	"github.com/nrc-no/notcore/internal/utils/pointers"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var foo = time.Time(time.Date(1992, 7, 31, 0, 0, 0, 0, time.UTC))

func TestUnmarshalIndividualsTabularData2(t *testing.T) {
	locales.LoadTranslations()
	locales.Init()
	tl := locales.GetTranslator()

	var tests = []struct {
		name           string
		data           [][]string
		colMapping     map[string]int
		rowLimit       *int
		expectedErrors []FileError
	}{
		{
			name:       "upload limit exceeded",
			data:       [][]string{{}, {}, {}, {}},
			colMapping: map[string]int{},
			rowLimit:   pointers.Int(2),
			expectedErrors: []FileError{
				{tl("error_upload_limit", 3, 2), nil},
			},
		},
		{
			name:           "id",
			data:           [][]string{{constants.DBColumnIndividualID}, {"unique-user-id"}},
			colMapping:     map[string]int{},
			rowLimit:       pointers.Int(200),
			expectedErrors: nil,
		},
		{
			name:       "birthdate invalid",
			data:       [][]string{{constants.DBColumnIndividualBirthDate}, {"31-07-1992"}},
			colMapping: map[string]int{constants.DBColumnIndividualBirthDate: 0},
			rowLimit:   pointers.Int(200),
			expectedErrors: []FileError{
				{tl("error_row_parse_fail", 2), []error{errors.New(tl("error_parse_birthdate_invalid", constants.DBColumnIndividualBirthDate, "31-07-1992", "parsing time \"31-07-1992\" as \"2006-01-02\": cannot parse \"7-1992\" as \"2006\""))}},
			},
		},
		{
			name:           "birthdate valid",
			data:           [][]string{{constants.DBColumnIndividualBirthDate}, {"1992-07-31"}},
			colMapping:     map[string]int{constants.DBColumnIndividualBirthDate: 0},
			rowLimit:       pointers.Int(200),
			expectedErrors: nil,
		},
		{
			name:       "birthdate before minimum",
			data:       [][]string{{constants.DBColumnIndividualBirthDate}, {"1892-07-31"}},
			colMapping: map[string]int{constants.DBColumnIndividualBirthDate: 0},
			rowLimit:   pointers.Int(200),
			expectedErrors: []FileError{
				{tl("error_row_parse_fail", 2), []error{errors.New(tl("error_parse_birthdate_minimum", constants.DBColumnIndividualBirthDate, "1892-07-31 00:00:00 +0000 UTC", "1900-01-01 00:00:00 +0000 UTC"))}},
			},
		},
		{
			name:           "age",
			data:           [][]string{{constants.DBColumnIndividualAge}, {"31"}},
			colMapping:     map[string]int{constants.DBColumnIndividualAge: 0},
			rowLimit:       pointers.Int(200),
			expectedErrors: nil,
		},
		{
			name:       "age invalid",
			data:       [][]string{{constants.DBColumnIndividualAge}, {"-31"}},
			colMapping: map[string]int{constants.DBColumnIndividualAge: 0},
			rowLimit:   pointers.Int(200),
			expectedErrors: []FileError{
				{tl("error_row_parse_fail", 2), []error{errors.New(tl("error_parse_age", constants.DBColumnIndividualAge, "-31"))}},
			},
		},
		{
			name:       "service",
			data:       [][]string{{constants.DBColumnIndividualServiceCC1}, {"notAService"}},
			colMapping: map[string]int{constants.DBColumnIndividualServiceCC1: 0},
			rowLimit:   pointers.Int(200),
			expectedErrors: []FileError{
				{tl("error_row_parse_fail", 2), []error{errors.New(tl("error_invalid_value_w_hint", constants.DBColumnIndividualServiceCC1, errors.New(tl("error_unknown_service_type", "notAService")), enumTypes.AllServiceCCs().String()))}},
			},
		},
		{
			name:           "is minor",
			data:           [][]string{{constants.DBColumnIndividualIsMinor}, {"no"}},
			colMapping:     map[string]int{constants.DBColumnIndividualIsMinor: 0},
			rowLimit:       pointers.Int(200),
			expectedErrors: nil,
		},
		{
			name:       "is minor, invalid",
			data:       [][]string{{constants.DBColumnIndividualIsMinor}, {"notAnOption"}},
			colMapping: map[string]int{constants.DBColumnIndividualIsMinor: 0},
			rowLimit:   pointers.Int(200),
			expectedErrors: []FileError{
				{tl("error_row_parse_fail", 2), []error{errors.New(tl("error_invalid_value_w_hint", constants.DBColumnIndividualIsMinor, errors.New(tl("error_unknown_optional_boolean", "notAnOption")), enumTypes.AllOptionalBooleans().String()))}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileErrors := UnmarshalIndividualsTabularData(tt.data, &[]*Individual{}, tt.colMapping, tt.rowLimit)
			assert.Equal(t, tt.expectedErrors, fileErrors)
		})
	}
}
