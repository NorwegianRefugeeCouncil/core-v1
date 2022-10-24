package server

import (
	"fmt"
	"net"
	"net/url"
	"regexp"
	"time"
)

type Options struct {
	Address              string
	DatabaseDriver       string
	DatabaseDSN          string
	LogoutURL            string
	LoginURL             string
	JwtGroupGlobalAdmin  string
	AuthHeaderName       string
	AuthHeaderFormat     string
	OIDCIssuerURL        string
	OAuthClientID        string
	TokenRefreshURL      string
	TokenRefreshInterval time.Duration
}

var globalAdminGroupRegex = regexp.MustCompile(`^[A-Za-z0-9_-]+(?: +[A-Za-z0-9_-]+)*$`)

func (o Options) validate() error {
	if len(o.Address) == 0 {
		return fmt.Errorf("address is required")
	}
	_, _, err := net.SplitHostPort(o.Address)
	if err != nil {
		return fmt.Errorf("address is invalid: %w", err)
	}
	if len(o.DatabaseDriver) == 0 {
		return fmt.Errorf("database driver is required")
	}
	switch o.DatabaseDriver {
	case "sqlite":
	case "postgres":
	default:
		return fmt.Errorf("database driver is invalid. must be one of: sqlite, postgres")
	}
	if len(o.DatabaseDSN) == 0 {
		return fmt.Errorf("database DSN is required")
	}
	if err := o.validateRequiredURLOption(o.LogoutURL, "Logout URL"); err != nil {
		return err
	}
	if err := o.validateRequiredURLOption(o.TokenRefreshURL, "Refresh URL"); err != nil {
		return err
	}
	if o.TokenRefreshInterval <= time.Minute {
		return fmt.Errorf("minimum value for token refresh duration is 1 minute")
	}
	if len(o.JwtGroupGlobalAdmin) == 0 {
		return fmt.Errorf("JWT group global admin is required")
	}
	if !globalAdminGroupRegex.MatchString(o.JwtGroupGlobalAdmin) {
		return fmt.Errorf("JWT group global admin is invalid")
	}
	if !isValidRFC7230HeaderName(o.AuthHeaderName) {
		return fmt.Errorf("auth header name is invalid")
	}
	if o.AuthHeaderFormat != AuthHeaderFormatJWT &&
		o.AuthHeaderFormat != AuthHeaderFormatBearerToken {
		return fmt.Errorf("auth header format is invalid. must be one of: %s, %s",
			AuthHeaderFormatJWT,
			AuthHeaderFormatBearerToken)

	}
	if err := o.validateRequiredURLOption(o.OIDCIssuerURL, "Issuer URL"); err != nil {
		return err
	}
	if len(o.OAuthClientID) == 0 {
		return fmt.Errorf("OAuth client ID is required")
	}
	return nil
}

func (o Options) validateRequiredURLOption(u string, name string) error {
	if len(u) == 0 {
		return fmt.Errorf("%s is required", name)
	}
	if _, err := url.Parse(u); err != nil {
		return fmt.Errorf("%s is invalid: %w", name, err)
	}
	return nil
}

// isValidRFC7230HeaderName reports whether s is a valid RFC 7230 header name.
// https://www.rfc-editor.org/rfc/rfc7230
func isValidRFC7230HeaderName(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if !rfc7230AllowedHeaderChars[c] {
			return false
		}
	}
	return true
}

var rfc7230AllowedHeaderChars = map[rune]bool{}

func init() {
	for _, c := range "!#$%&'*+-.^_`|~0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" {
		rfc7230AllowedHeaderChars[c] = true
	}
}
