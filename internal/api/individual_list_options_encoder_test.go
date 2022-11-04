package api

import (
	"html/template"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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
			want: "/countries/usa/individuals?name=fullName",
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
			o:    ListIndividualsOptions{CountryID: countryId, DisplacementStatuses: []string{"status"}},
			want: "/countries/usa/individuals?displacement_status=status",
		}, {
			name: "displacement status multiple",
			o:    ListIndividualsOptions{CountryID: countryId, DisplacementStatuses: []string{"status1", "status2"}},
			want: "/countries/usa/individuals?displacement_status=status1&displacement_status=status2",
		}, {
			name: "gender",
			o:    ListIndividualsOptions{CountryID: countryId, Genders: []string{"male"}},
			want: "/countries/usa/individuals?gender=male",
		}, {
			name: "gender multiple",
			o:    ListIndividualsOptions{CountryID: countryId, Genders: []string{"male", "female"}},
			want: "/countries/usa/individuals?gender=male&gender=female",
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
			o:    ListIndividualsOptions{CountryID: countryId, IDs: []string{"id1", "id2"}},
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
