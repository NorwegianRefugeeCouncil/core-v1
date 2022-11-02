package api

//go:generate go run ../../tools/codegen/main.go --type=Country --output=zzz_generated_country_fields.go
type Country struct {
	ID       string `json:"id" db:"id"`
	Code     string `json:"code" db:"code"`
	Name     string `json:"name" db:"name"`
	JwtGroup string `json:"jwtGroup" db:"jwt_group"`
}

type CountryList struct {
	Items []*Country `db:"items"`
}
