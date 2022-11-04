package api

import (
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/utils/pointers"
	"github.com/stretchr/testify/assert"
)

func TestNewIndividualListFromURLValues(t *testing.T) {

	now := time.Now()
	birthDate10YearsOld := calculateBirthDateFromAge(10, now)
	birthDate20YearsOld := calculateBirthDateFromAge(20, now)

	tests := []struct {
		name    string
		args    url.Values
		want    ListIndividualsOptions
		wantErr bool
	}{
		{
			name: "empty",
			args: url.Values{},
			want: ListIndividualsOptions{},
		}, {
			name: constants.FormParamsGetIndividualsSkip,
			args: url.Values{constants.FormParamsGetIndividualsSkip: []string{"1"}},
			want: ListIndividualsOptions{Skip: 1},
		}, {
			name:    "invalid skip",
			args:    url.Values{constants.FormParamsGetIndividualsSkip: []string{"abc"}},
			wantErr: true,
		}, {
			name:    "negative skip",
			args:    url.Values{constants.FormParamsGetIndividualsSkip: []string{"-10"}},
			wantErr: true,
		}, {
			name: constants.FormParamsGetIndividualsTake,
			args: url.Values{constants.FormParamsGetIndividualsTake: []string{"1"}},
			want: ListIndividualsOptions{Take: 1},
		}, {
			name:    "invalid take",
			args:    url.Values{constants.FormParamsGetIndividualsTake: []string{"abc"}},
			wantErr: true,
		}, {
			name:    "negative take",
			args:    url.Values{constants.FormParamsGetIndividualsTake: []string{"-10"}},
			wantErr: true,
		}, {
			name: constants.FormParamsGetIndividualsFullName,
			args: url.Values{constants.FormParamsGetIndividualsFullName: []string{"fullName"}},
			want: ListIndividualsOptions{FullName: "fullName"},
		}, {
			name: constants.FormParamsGetIndividualsAddress,
			args: url.Values{constants.FormParamsGetIndividualsAddress: []string{"address"}},
			want: ListIndividualsOptions{Address: "address"},
		}, {
			name: constants.FormParamsGetIndividualsEmail,
			args: url.Values{constants.FormParamsGetIndividualsEmail: []string{"email"}},
			want: ListIndividualsOptions{Email: "email"},
		}, {
			name: constants.FormParamsGetIndividualsPhoneNumber,
			args: url.Values{constants.FormParamsGetIndividualsPhoneNumber: []string{"phoneNumber"}},
			want: ListIndividualsOptions{PhoneNumber: "phoneNumber"},
		}, {
			name: constants.FormParamsGetIndividualsCountryID,
			args: url.Values{constants.FormParamsGetIndividualsCountryID: []string{"countryID"}},
			want: ListIndividualsOptions{CountryID: "countryID"},
		}, {
			name: constants.FormParamsGetIndividualsAgeFrom,
			args: url.Values{constants.FormParamsGetIndividualsAgeFrom: []string{strconv.Itoa(10)}},
			want: ListIndividualsOptions{BirthDateTo: &birthDate10YearsOld},
		}, {
			name:    "invalid age from",
			args:    url.Values{constants.FormParamsGetIndividualsAgeFrom: []string{"invalid"}},
			wantErr: true,
		}, {
			name: constants.FormParamsGetIndividualsAgeTo,
			args: url.Values{constants.FormParamsGetIndividualsAgeTo: []string{strconv.Itoa(20)}},
			want: ListIndividualsOptions{BirthDateFrom: &birthDate20YearsOld},
		}, {
			name:    "invalid age to",
			args:    url.Values{constants.FormParamsGetIndividualsAgeTo: []string{"invalid"}},
			wantErr: true,
		}, {
			name: constants.FormParamsGetIndividualsIsMinor,
			args: url.Values{constants.FormParamsGetIndividualsIsMinor: []string{"true"}},
			want: ListIndividualsOptions{IsMinor: pointers.Bool(true)},
		}, {
			name: constants.FormParamsGetIndividualsIsMinor,
			args: url.Values{constants.FormParamsGetIndividualsIsMinor: []string{"false"}},
			want: ListIndividualsOptions{IsMinor: pointers.Bool(false)},
		}, {
			name:    "invalid isMinor",
			args:    url.Values{constants.FormParamsGetIndividualsIsMinor: []string{"invalid"}},
			wantErr: true,
		}, {
			name: constants.FormParamsGetIndividualsPresentsProtectionConcerns,
			args: url.Values{constants.FormParamsGetIndividualsPresentsProtectionConcerns: []string{"true"}},
			want: ListIndividualsOptions{PresentsProtectionConcerns: pointers.Bool(true)},
		}, {
			name: constants.FormParamsGetIndividualsPresentsProtectionConcerns,
			args: url.Values{constants.FormParamsGetIndividualsPresentsProtectionConcerns: []string{"false"}},
			want: ListIndividualsOptions{PresentsProtectionConcerns: pointers.Bool(false)},
		}, {
			name:    "invalid presentsProtectionConcerns",
			args:    url.Values{constants.FormParamsGetIndividualsPresentsProtectionConcerns: []string{"invalid"}},
			wantErr: true,
		}, {
			name: constants.FormParamsGetIndividualsID,
			args: url.Values{constants.FormParamsGetIndividualsID: []string{"1"}},
			want: ListIndividualsOptions{IDs: []string{"1"}},
		}, {
			name: "multiple ids",
			args: url.Values{constants.FormParamsGetIndividualsID: []string{"1", "2", "2", "3"}},
			want: ListIndividualsOptions{IDs: []string{"1", "2", "3"}},
		}, {
			name: constants.FormParamsGetIndividualsDisplacementStatus,
			args: url.Values{constants.FormParamsGetIndividualsDisplacementStatus: []string{"idp", "idp", "refugee"}},
			want: ListIndividualsOptions{DisplacementStatuses: []string{"idp", "refugee"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ret ListIndividualsOptions
			err := NewIndividualListFromURLValues(tt.args, &ret)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			if !assert.NoError(t, err) {
				return
			}
			assert.Equal(t, tt.want, ret)
		})
	}

}
