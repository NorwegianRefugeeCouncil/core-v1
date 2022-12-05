package middleware

import (
	"reflect"
	"testing"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/containers"
)

func Test_parsePermissions(t *testing.T) {
	type args struct {
		allCountries     []*api.Country
		globalAdminGroup string
		userGroups       []string
		nrcOrganisation  string
	}
	tests := []struct {
		name string
		args args
		want *parsedPermissions
	}{
		{
			name: "global admin. no countries defined",
			args: args{
				allCountries:     []*api.Country{},
				globalAdminGroup: "global-admin",
				userGroups:       []string{"global-admin"},
				nrcOrganisation:  "NRC Uganda",
			},
			want: &parsedPermissions{
				isGlobalAdmin: true,
				countryIds:    containers.NewStringSet(),
			},
		}, {
			name: "global admin. with countries defined",
			args: args{
				allCountries: []*api.Country{
					{ID: "1", NrcOrganisations: "country 1"},
				},
				globalAdminGroup: "global-admin",
				userGroups:       []string{"global-admin"},
				nrcOrganisation:  "country 1",
			},
			want: &parsedPermissions{
				isGlobalAdmin: true,
				countryIds:    containers.NewStringSet("1"),
			},
		}, {
			name: "country access only",
			args: args{
				allCountries: []*api.Country{
					{ID: "1", NrcOrganisations: "country 1"},
					{ID: "2", NrcOrganisations: "country 2"},
				},
				globalAdminGroup: "global-admin",
				userGroups:       []string{},
				nrcOrganisation:  "country 1",
			},
			want: &parsedPermissions{
				isGlobalAdmin: false,
				countryIds:    containers.NewStringSet("1"),
			},
		}, {
			name: "country access only, multiple countries per nrc organisation",
			args: args{
				allCountries: []*api.Country{
					{ID: "1", NrcOrganisations: "country 1,country 3"},
				},
				userGroups:      []string{},
				nrcOrganisation: "country 3",
			},
			want: &parsedPermissions{
				isGlobalAdmin: false,
				countryIds:    containers.NewStringSet("1"),
			},
		}, {
			name: "country access only. no countries",
			args: args{
				allCountries:     []*api.Country{},
				globalAdminGroup: "global-admin",
				userGroups:       []string{},
				nrcOrganisation:  "NRC Uganda",
			},
			want: &parsedPermissions{
				isGlobalAdmin: false,
				countryIds:    containers.NewStringSet(),
			},
		}, {
			name: "country access only. no matching countries",
			args: args{
				allCountries: []*api.Country{
					{ID: "1", NrcOrganisations: "country 1"},
					{ID: "2", NrcOrganisations: "country 2"},
				},
				globalAdminGroup: "global-admin",
				userGroups:       []string{},
				nrcOrganisation:  "country 3",
			},
			want: &parsedPermissions{
				isGlobalAdmin: false,
				countryIds:    containers.NewStringSet(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parsePermissions(tt.args.allCountries, tt.args.globalAdminGroup, tt.args.userGroups, tt.args.nrcOrganisation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}
