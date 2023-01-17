package utils

import "github.com/nrc-no/notcore/internal/containers"

type JwtGroupOptions struct {
	GlobalAdmin string
	CanRead     string
	CanWrite    string
}

// ParsedPermissions is a helper struct to store the parsed permissions
type ParsedPermissions struct {
	IsGlobalAdmin bool
	CountryIds    containers.StringSet
	CanRead       bool
	CanWrite      bool
}
