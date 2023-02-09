package middleware

import (
	"github.com/nrc-no/notcore/internal/api/enumTypes"
	"github.com/nrc-no/notcore/internal/utils"
	"reflect"
	"testing"

	"github.com/nrc-no/notcore/internal/containers"
)

func Test_parsePermissions(t *testing.T) {
	jwtGroups := utils.JwtGroupOptions{
		GlobalAdmin: "global-admin",
		CanRead:     "can-read",
		CanWrite:    "can-write",
	}
	type args struct {
		allCountries    []*enumTypes.Country
		jwtGroups       utils.JwtGroupOptions
		userGroups      []string
		nrcOrganisation string
	}
	tests := []struct {
		name string
		args args
		want *ParsedPermissions
	}{
		{
			name: "global admin. no countries defined",
			args: args{
				allCountries:    []*enumTypes.Country{},
				jwtGroups:       jwtGroups,
				userGroups:      []string{"global-admin"},
				nrcOrganisation: "NRC HO",
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: true,
				CanWrite:      true,
				CanRead:       true,
				CountryIds:    containers.NewStringSet(),
			},
		}, {
			name: "global admin. with countries defined",
			args: args{
				allCountries: []*enumTypes.Country{
					{ID: "1", NrcOrganisations: containers.NewStringSet("NRC Country 1")},
				},
				jwtGroups:       jwtGroups,
				userGroups:      []string{"global-admin"},
				nrcOrganisation: "NRC Country 1",
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: true,
				CanWrite:      true,
				CanRead:       true,
				CountryIds:    containers.NewStringSet("1"),
			},
		}, {
			name: "country access only, no read or write permissions",
			args: args{
				allCountries: []*enumTypes.Country{
					{ID: "1", NrcOrganisations: containers.NewStringSet("NRC Country 1")},
					{ID: "2", NrcOrganisations: containers.NewStringSet("NRC Country 2")},
				},
				jwtGroups:       jwtGroups,
				userGroups:      []string{},
				nrcOrganisation: "NRC Country 1",
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: false,
				CanWrite:      false,
				CanRead:       false,
				CountryIds:    containers.NewStringSet("1"),
			},
		}, {
			name: "country access only, multiple countries per nrc organisation, no read or write permissions",
			args: args{
				allCountries: []*enumTypes.Country{
					{ID: "1", NrcOrganisations: containers.NewStringSet("NRC Country 1", "NRC Country 3")},
				},
				userGroups:      []string{},
				nrcOrganisation: "NRC Country 3",
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: false,
				CanWrite:      false,
				CanRead:       false,
				CountryIds:    containers.NewStringSet("1"),
			},
		}, {
			name: "country access only. no countries",
			args: args{
				allCountries:    []*enumTypes.Country{},
				jwtGroups:       jwtGroups,
				userGroups:      []string{},
				nrcOrganisation: "NRC Country 1",
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: false,
				CanWrite:      false,
				CanRead:       false,
				CountryIds:    containers.NewStringSet(),
			},
		}, {
			name: "country access only. no matching countries",
			args: args{
				allCountries: []*enumTypes.Country{
					{ID: "1", NrcOrganisations: containers.NewStringSet("NRC Country 1")},
					{ID: "2", NrcOrganisations: containers.NewStringSet("NRC Country 2")},
				},
				jwtGroups:       jwtGroups,
				userGroups:      []string{},
				nrcOrganisation: "NRC Country 3",
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: false,
				CanWrite:      false,
				CanRead:       false,
				CountryIds:    containers.NewStringSet(),
			},
		}, {
			name: "country access only, only read permissions",
			args: args{
				allCountries: []*enumTypes.Country{
					{ID: "1", NrcOrganisations: containers.NewStringSet("NRC Country 1")},
				},
				jwtGroups:       jwtGroups,
				userGroups:      []string{"can-read"},
				nrcOrganisation: "NRC Country 1",
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: false,
				CanWrite:      false,
				CanRead:       true,
				CountryIds:    containers.NewStringSet("1"),
			},
		}, {
			name: "country access only, only write permissions",
			args: args{
				allCountries: []*enumTypes.Country{
					{ID: "1", NrcOrganisations: containers.NewStringSet("NRC Country 1")},
				},
				jwtGroups:       jwtGroups,
				userGroups:      []string{"can-write"},
				nrcOrganisation: "NRC Country 1",
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: false,
				CanWrite:      true,
				CanRead:       true,
				CountryIds:    containers.NewStringSet("1"),
			},
		}, {
			name: "country access only, read and write permissions",
			args: args{
				allCountries: []*enumTypes.Country{
					{ID: "1", NrcOrganisations: containers.NewStringSet("NRC Country 1")},
				},
				jwtGroups:       jwtGroups,
				userGroups:      []string{"can-read", "can-write"},
				nrcOrganisation: "NRC Country 1",
			},
			want: &ParsedPermissions{
				IsGlobalAdmin: false,
				CanWrite:      true,
				CanRead:       true,
				CountryIds:    containers.NewStringSet("1"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parsePermissions(tt.args.allCountries, tt.args.jwtGroups, tt.args.userGroups, tt.args.nrcOrganisation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}
