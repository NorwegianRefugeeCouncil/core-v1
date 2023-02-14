package api

import "github.com/nrc-no/notcore/internal/containers"

type Country struct {
	ID               string               `db:"id"`
	Code             string               `db:"code"`
	Name             string               `db:"name"`
	NrcOrganisations containers.StringSet `db:"nrc_organisations"`
}

type CountryList struct {
	Items []*Country `db:"items"`
}
