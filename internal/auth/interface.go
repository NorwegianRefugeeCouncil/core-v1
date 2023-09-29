package auth

import (
	"github.com/nrc-no/notcore/internal/containers"
)

//go:generate mockgen -destination=./interface_mock.go -package=auth . Interface

type Permission uint8

type CountryPermissions map[string]containers.Set[Permission]

const (
	PermissionGlobalAdmin Permission = iota
	PermissionWrite
	PermissionRead
)

type Interface interface {
	IsGlobalAdmin() bool
	GetAllowedCountries() containers.StringSet
	HasCountryLevelPermission(countryID string, perm Permission) bool
	HasCountryPermissionWrite(countryID string) bool
	HasCountryPermissionRead(countryID string) bool
}

type permissions struct {
	// isGlobalAdmin is true if the user has global admin permissions
	isGlobalAdmin bool
	countryPermissions CountryPermissions
	allowedCountryIDs containers.StringSet
	// allCountryIDs is a list of all country IDs
	allCountryIDs containers.StringSet
}

func New(countryPermissions CountryPermissions, allCountryIDs containers.StringSet, isGlobalAdmin bool) Interface {
	allowedCountryIDs := containers.NewStringSet()
  for k := range countryPermissions{
		allowedCountryIDs.Add(k)
  }
	p := permissions{
		countryPermissions: countryPermissions,
		allowedCountryIDs: allowedCountryIDs,
		allCountryIDs:     containers.NewStringSet(allCountryIDs.Items()...),
		isGlobalAdmin:     isGlobalAdmin,
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
		return p.HasCountryPermissionWrite(countryID)
	case PermissionRead:
		return p.HasCountryPermissionRead(countryID)
	default:
		return false
	}
}

func (p permissions) GetAllowedCountries() containers.StringSet {
	if p.IsGlobalAdmin() {
		return containers.NewStringSet(p.allCountryIDs.Items()...)
	}
	return p.allowedCountryIDs 
}

func (p permissions) HasCountryPermissionWrite(countryID string) bool {
	return p.IsGlobalAdmin() || p.countryPermissions[countryID].Contains(PermissionWrite)
}

func (p permissions) HasCountryPermissionRead(countryID string) bool {
	return p.IsGlobalAdmin() || p.countryPermissions[countryID].Contains(PermissionRead) || p.countryPermissions[countryID].Contains(PermissionWrite) 
}

func (p permissions) IsGlobalAdmin() bool {
	return p.isGlobalAdmin
}
