package auth

import (
	"github.com/nrc-no/notcore/internal/containers"
)

type Interface interface {
	IsGlobalAdmin() bool
	CanReadWriteToCountryID(countryID string) bool
	GetCountryIDsWithReadWritePermissions() containers.StringSet
}

type permissions struct {
	// isGlobalAdmin is true if the user has global admin permissions
	isGlobalAdmin bool
	// allowedCountryIDs is a list of country IDs that the user has explicit read/write permissions to
	allowedCountryIDs containers.StringSet
	// allCountryIDs is a list of all country IDs
	allCountryIDs containers.StringSet
}

func New(allowedCountryIDs, allCountryIDs containers.StringSet, isGlobalAdmin bool) Interface {
	p := permissions{
		allowedCountryIDs: containers.NewStringSet(allowedCountryIDs.Items()...),
		allCountryIDs:     containers.NewStringSet(allCountryIDs.Items()...),
		isGlobalAdmin:     isGlobalAdmin,
	}
	return p
}

func (p permissions) IsGlobalAdmin() bool {
	return p.isGlobalAdmin
}

func (p permissions) CanReadWriteToCountryID(countryID string) bool {
	return p.IsGlobalAdmin() || p.hasExplicitReadWritePermissionInCountry(countryID)
}

func (p permissions) GetCountryIDsWithReadWritePermissions() containers.StringSet {
	if p.IsGlobalAdmin() {
		return containers.NewStringSet(p.allCountryIDs.Items()...)
	}
	return containers.NewStringSet(p.allowedCountryIDs.Items()...)
}

func (p permissions) hasExplicitReadWritePermissionInCountry(countryID string) bool {
	return p.allowedCountryIDs.Contains(countryID)
}
