package api

import (
	"html/template"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/containers"
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
			want: ListIndividualsOptions{IDs: containers.NewStringSet("1")},
		}, {
			name: "multiple ids",
			args: url.Values{constants.FormParamsGetIndividualsID: []string{"1", "2", "2", "3"}},
			want: ListIndividualsOptions{IDs: containers.NewStringSet("1", "2", "3")},
		}, {
			name: constants.FormParamsGetIndividualsDisplacementStatus,
			args: url.Values{constants.FormParamsGetIndividualsDisplacementStatus: []string{"idp", "idp", "refugee"}},
			want: ListIndividualsOptions{DisplacementStatuses: containers.NewSet[DisplacementStatus](DisplacementStatusIDP, DisplacementStatusRefugee)},
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

func TestListIndividualsOptions_QueryParams(t *testing.T) {
	now := time.Now()
	birthDate10YearsOld := calculateBirthDateFromAge(10, now)
	birthDate20YearsOld := calculateBirthDateFromAge(20, now)

	const countryId = "usa"
	tests := []struct {
		name string
		o    ListIndividualsOptions
		want string
	}{
		{
			name: "empty",
			o:    ListIndividualsOptions{CountryID: countryId},
			want: "/countries/usa/individuals",
		}, {
			name: "skip",
			o:    ListIndividualsOptions{CountryID: countryId, Skip: 1},
			want: "/countries/usa/individuals?skip=1",
		}, {
			name: "take",
			o:    ListIndividualsOptions{CountryID: countryId, Take: 1},
			want: "/countries/usa/individuals?take=1",
		}, {
			name: "fullName",
			o:    ListIndividualsOptions{CountryID: countryId, FullName: "fullName"},
			want: "/countries/usa/individuals?full_name=fullName",
		}, {
			name: "address",
			o:    ListIndividualsOptions{CountryID: countryId, Address: "address"},
			want: "/countries/usa/individuals?address=address",
		}, {
			name: "birthDateFrom",
			o:    ListIndividualsOptions{CountryID: countryId, BirthDateFrom: &birthDate10YearsOld},
			want: "/countries/usa/individuals?age_to=10",
		}, {
			name: "birthDateTo",
			o:    ListIndividualsOptions{CountryID: countryId, BirthDateTo: &birthDate20YearsOld},
			want: "/countries/usa/individuals?age_from=20",
		}, {
			name: "displacement status",
			o:    ListIndividualsOptions{CountryID: countryId, DisplacementStatuses: containers.NewSet[DisplacementStatus](DisplacementStatusIDP)},
			want: "/countries/usa/individuals?displacement_status=idp",
		}, {
			name: "displacement status multiple",
			o:    ListIndividualsOptions{CountryID: countryId, DisplacementStatuses: containers.NewSet[DisplacementStatus](DisplacementStatusIDP, DisplacementStatusRefugee)},
			want: "/countries/usa/individuals?displacement_status=idp&displacement_status=refugee",
		}, {
			name: "gender",
			o:    ListIndividualsOptions{CountryID: countryId, Genders: containers.NewSet[Gender]("male")},
			want: "/countries/usa/individuals?gender=male",
		}, {
			name: "gender multiple",
			o:    ListIndividualsOptions{CountryID: countryId, Genders: containers.NewSet[Gender](GenderMale, GenderFemale)},
			want: "/countries/usa/individuals?gender=female&gender=male",
		}, {
			name: "isMinor",
			o:    ListIndividualsOptions{CountryID: countryId, IsMinor: &[]bool{true}[0]},
			want: "/countries/usa/individuals?is_minor=true",
		}, {
			name: "isNotMinor",
			o:    ListIndividualsOptions{CountryID: countryId, IsMinor: &[]bool{false}[0]},
			want: "/countries/usa/individuals?is_minor=false",
		}, {
			name: "presentsProtectionConcerns",
			o:    ListIndividualsOptions{CountryID: countryId, PresentsProtectionConcerns: &[]bool{true}[0]},
			want: "/countries/usa/individuals?presents_protection_concerns=true",
		}, {
			name: "does not presentsProtectionConcerns",
			o:    ListIndividualsOptions{CountryID: countryId, PresentsProtectionConcerns: &[]bool{false}[0]},
			want: "/countries/usa/individuals?presents_protection_concerns=false",
		}, {
			name: "phoneNumber",
			o:    ListIndividualsOptions{CountryID: countryId, PhoneNumber: "phoneNumber"},
			want: "/countries/usa/individuals?phone_number=phoneNumber",
		}, {
			name: "email",
			o:    ListIndividualsOptions{CountryID: countryId, Email: "email"},
			want: "/countries/usa/individuals?email=email",
		}, {
			name: "address",
			o:    ListIndividualsOptions{CountryID: countryId, Address: "address"},
			want: "/countries/usa/individuals?address=address",
		}, {
			name: "ids",
			o:    ListIndividualsOptions{CountryID: countryId, IDs: containers.NewStringSet("id1", "id2")},
			want: "/countries/usa/individuals?id=id1&id=id2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.o.QueryParams()
			want := template.HTML(tt.want)
			assert.Equal(t, want, got)
		})
	}
}
