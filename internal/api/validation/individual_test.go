package validation

import (
	"testing"
	"time"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/pkg/api/validation"
	"github.com/stretchr/testify/assert"
)

type IndividualBuilder struct {
	individual *api.Individual
}

func NewIndividualBuilder() *IndividualBuilder {
	return &IndividualBuilder{
		individual: &api.Individual{},
	}
}

func (i *IndividualBuilder) Build() *api.Individual {
	return i.individual
}

func (i *IndividualBuilder) WithID(id string) *IndividualBuilder {
	i.individual.ID = id
	return i
}

func (i *IndividualBuilder) WithCountryID(countryID string) *IndividualBuilder {
	i.individual.CountryID = countryID
	return i
}

func (i *IndividualBuilder) WithFullName(fullName string) *IndividualBuilder {
	i.individual.FullName = fullName
	return i
}

func (i *IndividualBuilder) WithPreferredName(preferredName string) *IndividualBuilder {
	i.individual.PreferredName = preferredName
	return i
}

func (i *IndividualBuilder) WithEmail(email string) *IndividualBuilder {
	i.individual.Email = email
	return i
}

func (i *IndividualBuilder) WithAddress(address string) *IndividualBuilder {
	i.individual.Address = address
	return i
}

func (i *IndividualBuilder) WithBirthDate(birthDate *time.Time) *IndividualBuilder {
	i.individual.BirthDate = birthDate
	return i
}

func (i *IndividualBuilder) WithPhoneNumber(phoneNumber string) *IndividualBuilder {
	i.individual.PhoneNumber = phoneNumber
	return i
}

func (i *IndividualBuilder) WithDisplacementStatus(displacementStatus api.DisplacementStatus) *IndividualBuilder {
	i.individual.DisplacementStatus = displacementStatus
	return i
}

func (i *IndividualBuilder) WithGender(gender api.Gender) *IndividualBuilder {
	i.individual.Gender = gender
	return i
}

func (i *IndividualBuilder) WithIsMinor(isMinor bool) *IndividualBuilder {
	i.individual.IsMinor = isMinor
	return i
}

func (i *IndividualBuilder) WithPresentsProtectionConcerns(presentsProtectionConcerns bool) *IndividualBuilder {
	i.individual.PresentsProtectionConcerns = presentsProtectionConcerns
	return i
}

func (i *IndividualBuilder) WithPhysicalImpairment(physicalImpairment string) *IndividualBuilder {
	i.individual.PhysicalImpairment = physicalImpairment
	return i
}

func (i *IndividualBuilder) WithSensoryImpairment(sensoryImpairment string) *IndividualBuilder {
	i.individual.SensoryImpairment = sensoryImpairment
	return i
}

func (i *IndividualBuilder) WithMentalImpairment(metalImpairment string) *IndividualBuilder {
	i.individual.MentalImpairment = metalImpairment
	return i
}

func ValidIndividual() *IndividualBuilder {
	bd := time.Now().AddDate(-10, 0, 0)
	return NewIndividualBuilder().
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
			want: validation.ErrorList{validation.NotSupported(displacementStatusPath, api.DisplacementStatus("bla"), []string{"host_community", "idp", "refugee"})},
		}, {
			name: "empty gender",
			i:    ValidIndividual().WithGender("").Build(),
			want: validation.ErrorList{validation.Required(genderPath, "gender is required")},
		}, {
			name: "invalid gender",
			i:    ValidIndividual().WithGender("bla").Build(),
			want: validation.ErrorList{validation.NotSupported(genderPath, api.Gender("bla"), []string{"female", "male", "other", "prefers_not_to_say"})},
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
					api.DisplacementStatus("bla"),
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
