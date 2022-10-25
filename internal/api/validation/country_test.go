package validation

import (
	"strings"
	"testing"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/validation"
	"github.com/stretchr/testify/assert"
)

type CountryBuilder struct {
	country *api.Country
}

func ValidCountry() *CountryBuilder {
	return &CountryBuilder{
		country: &api.Country{
			ID:       "id",
			Code:     "code",
			Name:     "name",
			JwtGroup: "jwt_group",
		},
	}
}

func (b *CountryBuilder) Build() *api.Country {
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

func (b *CountryBuilder) WithJwtGroup(jwtGroup string) *CountryBuilder {
	b.country.JwtGroup = jwtGroup
	return b
}

func (b *CountryBuilder) WithID(id string) *CountryBuilder {
	b.country.ID = id
	return b
}

func TestValidateCountry(t *testing.T) {
	namePath := validation.NewPath("name")
	codePath := validation.NewPath("code")
	jwtGroupPath := validation.NewPath("jwtGroup")
	tests := []struct {
		name    string
		country *api.Country
		want    validation.ErrorList
	}{
		{
			name:    "valid",
			country: ValidCountry().Build(),
			want:    validation.ErrorList{},
		}, {
			name:    "invalid name",
			country: ValidCountry().WithName("!!").Build(),
			want:    validation.ErrorList{validation.Invalid(namePath, "!!", "country name can only contain letters or spaces")},
		}, {
			name:    "missing name",
			country: ValidCountry().WithName("").Build(),
			want:    validation.ErrorList{validation.Required(namePath, "country name is required")},
		}, {
			name:    "name too long",
			country: ValidCountry().WithName(longString(256)).Build(),
			want:    validation.ErrorList{validation.TooLongMaxLength(namePath, longString(256), 255)},
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
			country: ValidCountry().WithCode(longString(256)).Build(),
			want:    validation.ErrorList{validation.TooLongMaxLength(codePath, longString(256), 255)},
		}, {
			name:    "missing jwt group",
			country: ValidCountry().WithJwtGroup("").Build(),
			want:    validation.ErrorList{validation.Required(jwtGroupPath, "jwt group is required")},
		}, {
			name:    "invalid jwt group",
			country: ValidCountry().WithJwtGroup("!!").Build(),
			want:    validation.ErrorList{validation.Invalid(jwtGroupPath, "!!", "jwt group is invalid")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateCountry(tt.country)
			assert.Equal(t, tt.want, got)
		})
	}
}

func longString(n int) string {
	return strings.Repeat("a", n)
}
