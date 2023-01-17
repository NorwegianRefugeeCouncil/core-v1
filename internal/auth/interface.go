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
	GetAllowedCountries() containers.StringSet
	HasCountryLevelPermission(countryID string, perm Permission) bool
	HasMinimumPermissionWrite(countryID string) bool
	HasMinimumPermissionRead(countryID string) bool
	CanRead(countryID string) bool
	CanWrite(countryID string) bool
}

type permissions struct {
	// isGlobalAdmin is true if the user has global admin permissions
	isGlobalAdmin bool
	canRead       bool
	canWrite      bool
	// allowedCountryIDs is a list of country IDs that the user has explicit read/write permissions to
	allowedCountryIDs containers.StringSet
	// allCountryIDs is a list of all country IDs
	allCountryIDs containers.StringSet
}

func New(allowedCountryIDs, allCountryIDs containers.StringSet, isGlobalAdmin bool, canRead bool, canWrite bool) Interface {
	p := permissions{
		allowedCountryIDs: containers.NewStringSet(allowedCountryIDs.Items()...),
		allCountryIDs:     containers.NewStringSet(allCountryIDs.Items()...),
		isGlobalAdmin:     isGlobalAdmin,
		canRead:           canRead,
		canWrite:          canWrite,
	}
	return p
}

func (p permissions) HasCountryLevelPermission(countryID string, perm Permission) bool {
	if !p.allCountryIDs.Contains(countryID) {
		return p.IsGlobalAdmin()
	}

	switch perm {
	case PermissionGlobalAdmin:
		return p.IsGlobalAdmin()
	case PermissionWrite:
		return p.HasMinimumPermissionWrite(countryID)
	case PermissionRead:
		return p.HasMinimumPermissionRead(countryID)
	default:
		return false
	}
}

func (p permissions) GetAllowedCountries() containers.StringSet {
	if p.IsGlobalAdmin() {
		return containers.NewStringSet(p.allCountryIDs.Items()...)
	}
	return containers.NewStringSet(p.allowedCountryIDs.Items()...)
}

func (p permissions) HasMinimumPermissionWrite(countryID string) bool {
	return p.isGlobalAdmin || (p.allowedCountryIDs.Contains(countryID) && p.canWrite)
}

func (p permissions) HasMinimumPermissionRead(countryID string) bool {
	return p.isGlobalAdmin || (p.allowedCountryIDs.Contains(countryID) && (p.canWrite || p.canRead))
}

func (p permissions) CanWrite(countryID string) bool {
	return p.allowedCountryIDs.Contains(countryID) && p.canWrite
}

func (p permissions) CanRead(countryID string) bool {
	return p.allowedCountryIDs.Contains(countryID) && p.canRead
}

func (p permissions) IsGlobalAdmin() bool {
	return p.isGlobalAdmin
}
