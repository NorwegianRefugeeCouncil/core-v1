package validation

import (
	"testing"
	"time"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/pkg/api/validation"
	"github.com/stretchr/testify/assert"
)

func ValidIndividual() *api.Individual_Builder {
	bd := time.Now().AddDate(-10, 0, 0)
	return api.New_Individual_Builder().
		WithEmail("email@email.com").
		WithPhoneNumber("1234567890").
		WithFullName("John Doe").
		WithDisplacementStatus("idp").
		WithBirthDate(&bd).
		WithCountryID("countryID").
		WithPreferredName("John").
		WithGender("male")
}

func TestValidateIndividual(t *testing.T) {
	futureDate := time.Now().AddDate(1, 0, 0)
	emptyDate := time.Time{}
	emailPath := validation.NewPath("email")
	birthDatePath := validation.NewPath("birthDate")
	countryIDPath := validation.NewPath("countryId")
	displacementStatusPath := validation.NewPath("displacementStatus")
	genderPath := validation.NewPath("gender")
	tests := []struct {
		name string
		i    *api.Individual
		want validation.ErrorList
	}{
		{
			name: "valid",
			i:    ValidIndividual().Build(),
			want: validation.ErrorList{},
		}, {
			name: "invalid email",
			i:    ValidIndividual().WithEmail("invalid").Build(),
			want: validation.ErrorList{validation.Invalid(emailPath, "invalid", "invalid email address")},
		}, {
			name: "birth date in future",
			i:    ValidIndividual().WithBirthDate(&futureDate).Build(),
			want: validation.ErrorList{validation.Invalid(birthDatePath, &futureDate, "birthdate cannot be in the future")},
		}, {
			name: "birth date empty",
			i:    ValidIndividual().WithBirthDate(&emptyDate).Build(),
			want: validation.ErrorList{validation.Invalid(birthDatePath, &emptyDate, "must be a valid date")},
		}, {
			name: "empty country id",
			i:    ValidIndividual().WithCountryID("").Build(),
			want: validation.ErrorList{validation.Required(countryIDPath, "country id is required")},
		}, {
			name: "empty displacement status",
			i:    ValidIndividual().WithDisplacementStatus("").Build(),
			want: validation.ErrorList{validation.Required(displacementStatusPath, "displacement status is required")},
		}, {
			name: "invalid displacement status",
			i:    ValidIndividual().WithDisplacementStatus("bla").Build(),
			want: validation.ErrorList{validation.NotSupported(displacementStatusPath, "bla", []string{"host_community", "idp", "refugee"})},
		}, {
			name: "empty gender",
			i:    ValidIndividual().WithGender("").Build(),
			want: validation.ErrorList{validation.Required(genderPath, "gender is required")},
		}, {
			name: "invalid gender",
			i:    ValidIndividual().WithGender("bla").Build(),
			want: validation.ErrorList{validation.NotSupported(genderPath, "bla", []string{"female", "male", "other", "prefers_not_to_say"})},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateIndividual(tt.i)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestValidateIndividualList(t *testing.T) {
	tests := []struct {
		name string
		i    *api.IndividualList
		want validation.ErrorList
	}{
		{
			name: "valid",
			i: &api.IndividualList{
				Items: []*api.Individual{ValidIndividual().Build()},
			},
			want: validation.ErrorList{},
		}, {
			name: "invalid",
			i: &api.IndividualList{
				Items: []*api.Individual{ValidIndividual().WithDisplacementStatus("bla").Build()},
			},
			want: validation.ErrorList{
				validation.NotSupported(
					validation.NewPath("items[0].displacementStatus"),
					"bla",
					[]string{"host_community", "idp", "refugee"})},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateIndividualList(tt.i)
			assert.Equal(t, tt.want, got)
		})
	}
}
