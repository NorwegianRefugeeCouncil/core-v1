package auth

import (
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/containers"
)

type Interface interface {
	IsGlobalAdmin() bool
	CanWriteToCountryID(countryID string) bool
	CanReadFromCountryID(countryID string) bool
	CanAdminCountryID(countryID string) bool
	GetCountryIDsWithPermission(permission string) containers.StringSet
	HasExplicitPermission(countryID string, permission string) bool
}

type permissions struct {
	// countryMap from Country.ID -> Country
	countryMap map[string]*api.Country
	// countryCodeMap from Country.Code -> Country.ID
	countryCodeMap           map[string]string
	explicitReadPermissions  map[string]bool
	explicitWritePermissions map[string]bool
	explicitAdminPermissions map[string]bool
	isGlobalAdmin            bool
}

func New(userPermissions api.UserPermissions, countries []*api.Country) Interface {
	p := permissions{
		countryMap:               map[string]*api.Country{},
		countryCodeMap:           map[string]string{},
		explicitReadPermissions:  map[string]bool{},
		explicitWritePermissions: map[string]bool{},
		explicitAdminPermissions: map[string]bool{},
	}

	for _, country := range countries {
		p.countryMap[country.ID] = country
		p.countryCodeMap[country.Code] = country.ID
	}

	for countryID, permission := range userPermissions.ExplicitCountryPermissions {
		if permission.Read {
			p.explicitReadPermissions[countryID] = true
		}
		if permission.Write {
			p.explicitWritePermissions[countryID] = true
		}
		if permission.Admin {
			p.explicitAdminPermissions[countryID] = true
		}
	}
	p.isGlobalAdmin = userPermissions.IsGlobalAdmin

	return p
}

func (p permissions) IsGlobalAdmin() bool {
	return p.isGlobalAdmin
}

func (p permissions) CanWriteToCountryCode(code string) bool {
	if countryID, ok := p.getCountryID(code); ok {
		return p.CanWriteToCountryID(countryID)
	}
	return false
}

func (p permissions) HasExplicitPermission(countryID string, permission string) bool {
	if _, ok := p.countryMap[countryID]; !ok {
		return false
	}
	switch permission {
	case "read":
		return p.explicitReadPermissions[countryID]
	case "write":
		return p.explicitWritePermissions[countryID]
	case "admin":
		return p.explicitAdminPermissions[countryID]
	default:
		return false
	}
}

func (p permissions) CanWriteToCountryID(countryID string) bool {
	if _, ok := p.countryMap[countryID]; !ok {
		return false
	}
	return p.IsGlobalAdmin() || p.explicitWritePermissions[countryID] || p.explicitAdminPermissions[countryID]
}

func (p permissions) CanReadFromCountryID(countryID string) bool {
	if _, ok := p.countryMap[countryID]; !ok {
		return false
	}
	return p.IsGlobalAdmin() || p.explicitReadPermissions[countryID] || p.explicitWritePermissions[countryID] || p.explicitAdminPermissions[countryID]
}

func (p permissions) CanAdminCountryID(countryID string) bool {
	if _, ok := p.countryMap[countryID]; !ok {
		return false
	}
	return p.IsGlobalAdmin() || p.explicitAdminPermissions[countryID]
}

func (p permissions) GetCountryIDsWithPermission(permission string) containers.StringSet {
	return p.getCountryIDsWithAnyPermission(permission)
}

func (p permissions) getCountryID(code string) (string, bool) {
	countryID, ok := p.countryCodeMap[code]
	return countryID, ok
}

func (p permissions) getCountryIDsWithAnyPermission(permissions ...string) containers.StringSet {
	ret := containers.NewStringSet()
	countries := p.getCountriesWithAnyPermission(permissions...)
	for _, c := range countries {
		ret.Add(c.ID)
	}
	return ret
}

func (p permissions) getCountriesWithAnyPermission(permissions ...string) []*api.Country {
	var ret []*api.Country
	if p.IsGlobalAdmin() {
		for _, c := range p.countryMap {
			ret = append(ret, c)
		}
		return ret
	}
	ids := containers.NewStringSet()
	for _, permission := range permissions {
		switch permission {
		case "read":
			for countryID := range p.explicitReadPermissions {
				ids.Add(countryID)
			}
			for countryID := range p.explicitWritePermissions {
				ids.Add(countryID)
			}
			for countryID := range p.explicitAdminPermissions {
				ids.Add(countryID)
			}
		case "write":
			for countryID := range p.explicitWritePermissions {
				ids.Add(countryID)
			}
			for countryID := range p.explicitAdminPermissions {
				ids.Add(countryID)
			}
		case "admin":
			for countryID := range p.explicitAdminPermissions {
				ids.Add(countryID)
			}
		default:
			return nil
		}
	}
	for _, id := range ids.Items() {
		ret = append(ret, p.countryMap[id])
	}
	return ret
}
