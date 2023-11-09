package api

type Country struct {
	ID               string               `db:"id"`
	Code             string               `db:"code"`
	Name             string               `db:"name"`
	ReadGroup 			 string               `db:"read_group"`
	WriteGroup 			 string               `db:"write_group"`
}

type CountryList struct {
	Items []*Country `db:"items"`
}
