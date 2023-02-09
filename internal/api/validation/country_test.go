package validation

import (
	"github.com/nrc-no/notcore/internal/api/enumTypes"
	"github.com/nrc-no/notcore/internal/containers"
	"strings"
	"testing"

	"github.com/nrc-no/notcore/pkg/api/validation"
	"github.com/stretchr/testify/assert"
)

type CountryBuilder struct {
	country *enumTypes.Country
}

func ValidCountry() *CountryBuilder {
	return &CountryBuilder{
		country: &enumTypes.Country{
			ID:               "id",
			Code:             "code",
			Name:             "name",
			NrcOrganisations: containers.NewStringSet("nrc_organisation"),
		},
	}
}

func (b *CountryBuilder) Build() *enumTypes.Country {
	return b.country
}

func (b *CountryBuilder) WithName(name string) *CountryBuilder {
	b.country.Name = name
	return b
}

func (b *CountryBuilder) WithCode(code string) *CountryBuilder {
	b.country.Code = code
	return b
}

func (b *CountryBuilder) WithNrcOrganisation(nrcOrganisations containers.StringSet) *CountryBuilder {
	b.country.NrcOrganisations = nrcOrganisations
	return b
}

func (b *CountryBuilder) WithID(id string) *CountryBuilder {
	b.country.ID = id
	return b
}

func TestValidateCountry(t *testing.T) {
	namePath := validation.NewPath("name")
	codePath := validation.NewPath("code")
	nrcOrganisationPath := validation.NewPath("nrcOrganisations")
	weirdString := string([]byte{0x7f, 0x7f})
	tests := []struct {
		name    string
		country *enumTypes.Country
		want    validation.ErrorList
	}{
		{
			name:    "valid",
			country: ValidCountry().Build(),
			want:    validation.ErrorList{},
		}, {
			name:    "invalid name",
			country: ValidCountry().WithName(weirdString).Build(),
			want:    validation.ErrorList{validation.Invalid(namePath, weirdString, "country name can only contain letters or spaces")},
		}, {
			name:    "missing name",
			country: ValidCountry().WithName("").Build(),
			want:    validation.ErrorList{validation.Required(namePath, "country name is required")},
		}, {
			name:    "name too long",
			country: ValidCountry().WithName(bigstr(256)).Build(),
			want:    validation.ErrorList{validation.TooLongMaxLength(namePath, bigstr(256), 255)},
		}, {
			name:    "name too short",
			country: ValidCountry().WithName("a").Build(),
			want:    validation.ErrorList{validation.TooShortMinLength(namePath, "a", 2)},
		}, {
			name:    "missing code",
			country: ValidCountry().WithCode("").Build(),
			want:    validation.ErrorList{validation.Required(codePath, "country code is required")},
		}, {
			name:    "invalid code",
			country: ValidCountry().WithCode("!!").Build(),
			want:    validation.ErrorList{validation.Invalid(codePath, "!!", "country code can only contain letters, numbers, underscores and hyphens")},
		}, {
			name:    "code too short",
			country: ValidCountry().WithCode("a").Build(),
			want:    validation.ErrorList{validation.TooShortMinLength(codePath, "a", 2)},
		}, {
			name:    "code too long",
			country: ValidCountry().WithCode(bigstr(256)).Build(),
			want:    validation.ErrorList{validation.TooLongMaxLength(codePath, bigstr(256), 255)},
		}, {
			name:    "missing nrc organisation",
			country: ValidCountry().WithNrcOrganisation(containers.NewStringSet()).Build(),
			want:    validation.ErrorList{validation.Required(nrcOrganisationPath, "nrc organisation is required")},
		}, {
			name:    "invalid nrc organisation",
			country: ValidCountry().WithNrcOrganisation(containers.NewStringSet("!!")).Build(),
			want:    validation.ErrorList{validation.Invalid(nrcOrganisationPath, "!!", "nrc organisation is invalid")},
		}, {
			name:    "nrc organisation too long",
			country: ValidCountry().WithNrcOrganisation(containers.NewStringSet(bigstr(256), "shortName")).Build(),
			want:    validation.ErrorList{validation.TooLongMaxLength(nrcOrganisationPath, bigstr(256), 255)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateCountry(tt.country)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestValidateCountryList(t *testing.T) {
	tests := []struct {
		name        string
		countryList *enumTypes.CountryList
		want        validation.ErrorList
	}{
		{
			name:        "empty",
			countryList: &enumTypes.CountryList{},
			want:        validation.ErrorList{},
		}, {
			name: "valid",
			countryList: &enumTypes.CountryList{
				Items: []*enumTypes.Country{
					ValidCountry().Build(),
				},
			},
			want: validation.ErrorList{},
		}, {
			name: "invalid",
			countryList: &enumTypes.CountryList{
				Items: []*enumTypes.Country{
					ValidCountry().WithName("").Build(),
				},
			},
			want: validation.ErrorList{
				validation.Required(validation.NewPath("items[0].name"), "country name is required"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateCountryList(tt.countryList)
			assert.Equal(t, tt.want, got)
		})
	}
}

func bigstr(n int) string {
	return strings.Repeat("a", n)
}
