package api

type Country struct {
	ID              string `db:"id"`
	Code            string `db:"code"`
	Name            string `db:"name"`
	JwtGroup        string `db:"jwt_group"`
	NrcOrganisation string `db:"nrc_organisation"`
}

type CountryList struct {
	Items []*Country `db:"items"`
}
