package api

type Permission string

type CountryPermission struct {
	CountryID string
	Read      bool
	Write     bool
	Admin     bool
}

func (c CountryPermission) HasPermission(permission string) bool {
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
	IsGlobalAdmin      bool
	CountryPermissions CountryPermissions
}

func (u UserPermissions) HasCountryPermission(countryID, permission string) bool {
	if u.CountryPermissions == nil {
		return false
	}
	return u.IsGlobalAdmin || u.CountryPermissions.HasPermission(countryID, permission)
}

//func (u UserPermissions) HasCountryPermission(countryID, permission string) bool {
//	if u.CountryPermissions == nil {
//		return false
//	}
//	return u.IsGlobalAdmin || u.CountryPermissions.HasPermission(countryID, permission)
//}

func (u UserPermissions) GetCountryIDsWithPermission(permission string) []string {
	if u.CountryPermissions == nil {
		return []string{}
	}
	strMap := make(map[string]bool)
	for countryID, countryPermission := range u.CountryPermissions {
		if countryPermission.HasPermission(permission) {
			strMap[countryID] = true
		}
	}
	var ret []string
	for countryID := range strMap {
		ret = append(ret, countryID)
	}
	return ret
}

func (u UserPermissions) GetCountryIDsWithAnyPermission(permissions ...string) []string {
	if u.CountryPermissions == nil {
		return []string{}
	}
	strMap := make(map[string]bool)
	for countryID, countryPermission := range u.CountryPermissions {
		for _, permission := range permissions {
			if countryPermission.HasPermission(permission) {
				strMap[countryID] = true
				break
			}
		}
	}
	var ret []string
	for countryID := range strMap {
		ret = append(ret, countryID)
	}
	return ret
}
