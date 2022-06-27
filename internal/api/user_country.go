package api

type Permission string

type CountryPermission struct {
	CountryID string
	Read      bool
	Write     bool
	Admin     bool
}

type CountryPermissions map[string]CountryPermission

func (c CountryPermissions) HasPermission(countryID string, permission string) bool {
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
	UserID             string
	CountryPermissions CountryPermissions
}
