package api

type Permission string

type ExplicitCountryPermission struct {
	CountryID string
	Read      bool
	Write     bool
	Admin     bool
}

func (c ExplicitCountryPermission) HasPermission(permission string) bool {
	switch permission {
	case "read":
		return c.Read
	case "write":
		return c.Write
	case "admin":
		return c.Admin
	default:
		return false
	}
}

type ExplicitCountryPermissions map[string]ExplicitCountryPermission

func (c ExplicitCountryPermissions) Has(countryID string, permission string) bool {
	if c == nil {
		return false
	}
	if _, ok := c[countryID]; !ok {
		return false
	}
	switch permission {
	case "read":
		return c[countryID].Read
	case "write":
		return c[countryID].Write
	case "admin":
		return c[countryID].Admin
	default:
		return false
	}
}

type UserPermissions struct {
	UserID                     string
	IsGlobalAdmin              bool
	ExplicitCountryPermissions ExplicitCountryPermissions
}
