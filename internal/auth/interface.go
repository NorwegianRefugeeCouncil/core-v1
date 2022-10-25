package auth

import (
	"github.com/nrc-no/notcore/internal/containers"
)

//go:generate mockgen -destination=./interface_mock.go -package=auth . Interface

type Permission uint8

const (
	PermissionGlobalAdmin Permission = iota
	PermissionWrite
	PermissionRead
)

type Interface interface {
	IsGlobalAdmin() bool
	GetCountryIDsWithReadWritePermissions() containers.StringSet
	HasCountryLevelPermission(countryID string, perm Permission) bool
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

func (p permissions) HasCountryLevelPermission(countryID string, perm Permission) bool {
	if !p.allCountryIDs.Contains(countryID) {
		return false
	}

	switch perm {
	case PermissionGlobalAdmin:
		return p.IsGlobalAdmin()
	case PermissionWrite:
		return p.IsGlobalAdmin() || p.hasExplicitReadWritePermissionInCountry(countryID)
	case PermissionRead:
		return p.IsGlobalAdmin() || p.hasExplicitReadWritePermissionInCountry(countryID)
	default:
		return false
	}
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

func (p permissions) countryExists(countryID string) bool {
	return p.allCountryIDs.Contains(countryID)
}
