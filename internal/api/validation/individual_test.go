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

func (i *IndividualBuilder) WithEmail1(email string) *IndividualBuilder {
	i.individual.Email1 = email
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

func (i *IndividualBuilder) WithPhoneNumber1(phoneNumber string) *IndividualBuilder {
	i.individual.PhoneNumber1 = phoneNumber
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

func (i *IndividualBuilder) WithPreferredContactMethod(preferredContactMethod string) *IndividualBuilder {
	i.individual.PreferredContactMethod = preferredContactMethod
	return i
}

func (i *IndividualBuilder) WithCollectionAgentName(collectionAgentName string) *IndividualBuilder {
	i.individual.CollectionAgentName = collectionAgentName
	return i
}

func (i *IndividualBuilder) WithCollectionAgentTitle(collectionAgentTitle string) *IndividualBuilder {
	i.individual.CollectionAgentTitle = collectionAgentTitle
	return i
}

func (i *IndividualBuilder) WithCollectionDate(collectionDate time.Time) *IndividualBuilder {
	i.individual.CollectionTime = collectionDate
	return i
}

func ValidIndividual() *IndividualBuilder {
	bd := time.Now().AddDate(-10, 0, 0)
	return NewIndividualBuilder().
		WithEmail1("email@email.com").
		WithPhoneNumber1("1234567890").
		WithFullName("John Doe").
		WithDisplacementStatus("idp").
		WithBirthDate(&bd).
		WithCountryID("countryID").
		WithPreferredName("John").
		WithGender("male").
		WithCollectionAgentTitle("Collection Agent Title").
		WithCollectionAgentName("Collection Agent Name").
		WithCollectionDate(time.Now()).
		WithPreferredContactMethod("email")
}

func TestValidateIndividual(t *testing.T) {
	futureDate := time.Now().AddDate(1, 0, 0)
	emptyDate := time.Time{}
	email1Path := validation.NewPath("email1")
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
			name: "invalid email 1",
			i:    ValidIndividual().WithEmail1("invalid").Build(),
			want: validation.ErrorList{validation.Invalid(email1Path, "invalid", "invalid email address")},
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
			name: "invalid displacement status",
			i:    ValidIndividual().WithDisplacementStatus("bla").Build(),
			want: validation.ErrorList{validation.NotSupported(displacementStatusPath, api.DisplacementStatus("bla"), allowedDisplacementStatusesStr)},
		}, {
			name: "invalid gender",
			i:    ValidIndividual().WithGender("bla").Build(),
			want: validation.ErrorList{validation.NotSupported(genderPath, api.Gender("bla"), allowedGendersStr)},
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
					allowedDisplacementStatusesStr)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateIndividualList(tt.i)
			assert.Equal(t, tt.want, got)
		})
	}
}
