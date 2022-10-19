package server

import (
	"testing"
)

func (o Options) WithAddress(address string) Options {
	o.Address = address
	return o
}

func (o Options) WithDatabaseDriver(databaseDriver string) Options {
	o.DatabaseDriver = databaseDriver
	return o
}

func (o Options) WithDatabaseDSN(databaseDSN string) Options {
	o.DatabaseDSN = databaseDSN
	return o
}

func (o Options) WithLogoutURL(logoutURL string) Options {
	o.LogoutURL = logoutURL
	return o
}

func (o Options) WithJwtGroupGlobalAdmin(jwtGroupGlobalAdmin string) Options {
	o.JwtGroupGlobalAdmin = jwtGroupGlobalAdmin
	return o
}

func (o Options) WithIDTokenHeaderName(idTokenHeaderName string) Options {
	o.IDTokenHeaderName = idTokenHeaderName
	return o
}

func validOptions() Options {
	return Options{
		Address:             ":8080",
		DatabaseDriver:      "postgres",
		DatabaseDSN:         "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
		LogoutURL:           "http://localhost:8080",
		JwtGroupGlobalAdmin: "global-admin",
		IDTokenHeaderName:   "X-Auth-Token",
	}
}

func TestOptions_validate(t *testing.T) {
	tests := []struct {
		name    string
		options Options
		wantErr bool
	}{
		{
			name:    "valid",
			options: validOptions(),
			wantErr: false,
		},
		{
			name:    "address is required",
			options: validOptions().WithAddress(""),
			wantErr: true,
		},
		{
			name:    "address is invalid",
			options: validOptions().WithAddress("invalid"),
			wantErr: true,
		},
		{
			name:    "database driver is required",
			options: validOptions().WithDatabaseDriver(""),
			wantErr: true,
		},
		{
			name:    "database driver is invalid",
			options: validOptions().WithDatabaseDriver("invalid"),
			wantErr: true,
		},
		{
			name:    "database DSN is required",
			options: validOptions().WithDatabaseDSN(""),
			wantErr: true,
		},
		{
			name:    "logout URL is required",
			options: validOptions().WithLogoutURL(""),
			wantErr: true,
		},
		{
			name:    "logout URL is invalid",
			options: validOptions().WithLogoutURL(string([]byte{0x7f})),
			wantErr: true,
		},
		{
			name:    "JWT group global admin is required",
			options: validOptions().WithJwtGroupGlobalAdmin(""),
			wantErr: true,
		},
		{
			name:    "JWT group global admin is invalid",
			options: validOptions().WithJwtGroupGlobalAdmin("!!!"),
			wantErr: true,
		},
		{
			name:    "ID token header name is required",
			options: validOptions().WithIDTokenHeaderName(""),
			wantErr: true,
		},
		{
			name:    "ID token header name is invalid",
			options: validOptions().WithIDTokenHeaderName("   "),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.options.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGlobalAdminGroupRegex(t *testing.T) {
	tests := []struct {
		name  string
		group string
		want  bool
	}{
		{
			name:  "valid",
			group: "globaladmin",
			want:  true,
		},
		{
			name:  "valid with dash",
			group: "global-admin",
			want:  true,
		},
		{
			name:  "valid with uppercase",
			group: "GlobalAdmin",
			want:  true,
		},
		{
			name:  "valid with underscores",
			group: "global_admin",
			want:  true,
		},
		{
			name:  "invalid with leading space",
			group: " global-admin",
			want:  false,
		},
		{
			name:  "invalid with trailing space",
			group: "global-admin ",
			want:  false,
		},
		{
			name:  "invalid with leading and trailing space",
			group: " global-admin ",
			want:  false,
		},
		{
			name:  "invalid with accented character",
			group: "è-Admin",
			want:  false,
		},
		{
			name:  "invalid with special character",
			group: "global!admin",
			want:  false,
		},
		{
			name:  "invalid with emoji",
			group: "👺",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := globalAdminGroupRegex.MatchString(tt.group); got != tt.want {
				t.Errorf("globalAdminGroupRegex.MatchString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidRFC7230HeaderName(t *testing.T) {
	tests := []struct {
		name   string
		header string
		want   bool
	}{
		{
			name:   "valid",
			header: "xauthtoken",
			want:   true,
		},
		{
			name:   "valid with dash",
			header: "x-auth-token",
			want:   true,
		},
		{
			name:   "valid with uppercase",
			header: "X-Auth-Token",
			want:   true,
		},
		{
			name:   "valid with underscores",
			header: "x_auth_token",
			want:   true,
		},
		{
			name:   "invalid with leading space",
			header: " x-auth-token",
			want:   false,
		},
		{
			name:   "invalid with trailing space",
			header: "x-auth-token ",
			want:   false,
		},
		{
			name:   "invalid with leading and trailing space",
			header: " x-auth-token ",
			want:   false,
		},
		{
			name:   "invalid with accented character",
			header: "è-Auth-Token",
			want:   false,
		},
		{
			name:   "invalid with emoji",
			header: "👺",
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidRFC7230HeaderName(tt.header); got != tt.want {
				t.Errorf("idTokenHeaderNameRegex.MatchString() = %v, want %v", got, tt.want)
			}
		})
	}
}
