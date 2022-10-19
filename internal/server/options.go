package server

import (
	"fmt"
	"net"
	"net/url"
	"regexp"
)

type Options struct {
	Address             string
	DatabaseDriver      string
	DatabaseDSN         string
	LogoutURL           string
	JwtGroupGlobalAdmin string
	IDTokenHeaderName   string
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
	if len(o.LogoutURL) == 0 {
		return fmt.Errorf("logout URL is required")
	}
	if _, err := url.Parse(o.LogoutURL); err != nil {
		return fmt.Errorf("logout URL is invalid: %w", err)
	}
	if len(o.JwtGroupGlobalAdmin) == 0 {
		return fmt.Errorf("JWT group global admin is required")
	}
	if !globalAdminGroupRegex.MatchString(o.JwtGroupGlobalAdmin) {
		return fmt.Errorf("JWT group global admin is invalid")
	}
	if !isValidRFC7230HeaderName(o.IDTokenHeaderName) {
		return fmt.Errorf("ID token header name is invalid")
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
