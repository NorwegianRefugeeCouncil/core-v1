package handlers

import (
	"context"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/utils"
)

type permissionHelper struct {
	countryMap       map[string]*api.Country
	countryCodeMap   map[string]string
	readPermissions  map[string]bool
	writePermissions map[string]bool
	adminPermissions map[string]bool
	isGlobalAdmin    bool
}

func newPermissionHelper(ctx context.Context, countries []*api.Country) permissionHelper {
	p := permissionHelper{
		countryMap:       map[string]*api.Country{},
		countryCodeMap:   map[string]string{},
		readPermissions:  map[string]bool{},
		writePermissions: map[string]bool{},
		adminPermissions: map[string]bool{},
	}

	for _, country := range countries {
		p.countryMap[country.ID] = country
		p.countryCodeMap[country.Code] = country.ID
	}

	perms := utils.GetRequestUserPermissions(ctx)
	for countryID, permission := range perms.CountryPermissions {
		if permission.Read {
			p.readPermissions[countryID] = true
		}
		if permission.Write {
			p.writePermissions[countryID] = true
		}
		if permission.Admin {
			p.adminPermissions[countryID] = true
		}
	}
	p.isGlobalAdmin = perms.IsGlobalAdmin

	return p
}

func (p permissionHelper) IsGlobalAdmin() bool {
	return p.isGlobalAdmin
}

func (p permissionHelper) GetCountryID(code string) (string, bool) {
	countryID, ok := p.countryCodeMap[code]
	return countryID, ok
}

func (p permissionHelper) HasCountryID(countryID string) bool {
	_, ok := p.countryCodeMap[countryID]
	return ok
}

func (p permissionHelper) HasCountryCode(code string) bool {
	_, ok := p.countryCodeMap[code]
	return ok
}

func (p permissionHelper) GetCountryByCode(code string) *api.Country {
	countryID, ok := p.countryCodeMap[code]
	if !ok {
		return nil
	}
	return p.countryMap[countryID]
}

func (p permissionHelper) CanWriteToCountryCode(code string) bool {
	if countryID, ok := p.GetCountryID(code); ok {
		return p.CanWriteToCountryID(countryID)
	}
	return false
}

func (p permissionHelper) CanReadFromCountryCode(code string) bool {
	if countryID, ok := p.GetCountryID(code); ok {
		return p.CanReadFromCountryID(countryID)
	}
	return false
}

func (p permissionHelper) CanAdminCountryCode(code string) bool {
	if countryID, ok := p.GetCountryID(code); ok {
		return p.CanAdminCountryID(countryID)
	}
	return false
}

func (p permissionHelper) CanWriteToCountryID(countryID string) bool {
	return p.writePermissions[countryID]
}

func (p permissionHelper) CanReadFromCountryID(countryID string) bool {
	return p.readPermissions[countryID]
}

func (p permissionHelper) CanAdminCountryID(countryID string) bool {
	return p.adminPermissions[countryID]
}

func (p permissionHelper) CanReadIndividual(individual *api.Individual) bool {

	if individual == nil {
		return false
	}
	if p.isGlobalAdmin {
		return true
	}
	for _, countryCode := range individual.Countries {
		if p.CanReadFromCountryCode(countryCode) {
			return true
		}
	}
	return false
}

func (p permissionHelper) GetCountryIDsWithPermission(permission string) containers.Set[string] {
	ret := containers.NewStringSet()

	if permission == "read" {
		for countryID, hasReadPermission := range p.readPermissions {
			if hasReadPermission {
				ret.Add(countryID)
			}
		}
	} else if permission == "write" {
		for countryID, hasWritePermission := range p.writePermissions {
			if hasWritePermission {
				ret.Add(countryID)
			}
		}
	} else if permission == "admin" {
		for countryID, hasAdminPermission := range p.adminPermissions {
			if hasAdminPermission {
				ret.Add(countryID)
			}
		}
	} else {
		return containers.NewStringSet()
	}
	return ret
}

func (p permissionHelper) GetCountryCodesWithPermission(permission string) containers.Set[string] {
	ret := containers.NewStringSet()
	switch permission {
	case "read":
		for countryID := range p.readPermissions {
			ret.Add(p.countryMap[countryID].Code)
		}
	case "write":
		for countryID := range p.writePermissions {
			ret.Add(p.countryMap[countryID].Code)
		}
	case "admin":
		for countryID := range p.adminPermissions {
			ret.Add(p.countryMap[countryID].Code)
		}
	default:
		return nil
	}
	return ret
}

func (p permissionHelper) GetCountryCodesWithAnyPermission(permissions ...string) containers.Set[string] {
	ret := containers.NewStringSet()
	for _, permission := range permissions {
		switch permission {
		case "read":
			for countryID := range p.readPermissions {
				ret.Add(p.countryMap[countryID].Code)
			}
		case "write":
			for countryID := range p.writePermissions {
				ret.Add(p.countryMap[countryID].Code)
			}
		case "admin":
			for countryID := range p.adminPermissions {
				ret.Add(p.countryMap[countryID].Code)
			}
		default:
			return nil
		}
	}
	return ret
}
