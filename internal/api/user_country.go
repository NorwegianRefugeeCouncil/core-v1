package api

type Permission string

const (
	ReadPermission  Permission = "read"
	WritePermission Permission = "write"
	AdminPermission Permission = "admin"
)

type UserCountry struct {
	UserID     string     `db:"user_id"`
	CountryID  string     `db:"country_id"`
	Permission Permission `db:"permission"`
}

type CountryPermission struct {
	CountryID  string
	Permission Permission
}

type UserPermissions struct {
	UserID             string
	CountryPermissions []CountryPermission
}

func (u UserPermissions) ToUserCountryList() []UserCountry {
	var ret []UserCountry
	for _, c := range u.CountryPermissions {
		ret = append(ret, UserCountry{
			UserID:     u.UserID,
			CountryID:  c.CountryID,
			Permission: c.Permission,
		})
	}
	return ret
}
