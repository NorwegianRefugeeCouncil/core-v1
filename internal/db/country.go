package db

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/rs/xid"
)

type CountryRepo interface {
	GetAll(ctx context.Context) ([]*api.Country, error)
	GetByID(ctx context.Context, id string) (*api.Country, error)
	Put(ctx context.Context, country *api.Country) (*api.Country, error)
}

type countryRepo struct {
	db *sqlx.DB
}

func NewCountryRepo(db *sqlx.DB) CountryRepo {
	return &countryRepo{db: db}
}

func (c countryRepo) GetAll(ctx context.Context) ([]*api.Country, error) {
	var countries []*api.Country
	err := c.db.SelectContext(ctx, &countries, "SELECT * FROM countries ORDER BY name")
	if err != nil {
		return nil, err
	}
	return countries, nil
}

func (c countryRepo) GetByID(ctx context.Context, id string) (*api.Country, error) {
	var country api.Country
	err := c.db.GetContext(ctx, &country, "SELECT * FROM countries WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &country, nil
}

func (c countryRepo) Put(ctx context.Context, country *api.Country) (*api.Country, error) {
	if country.ID == "" {
		country.ID = xid.New().String()
		_, err := c.db.ExecContext(ctx, "INSERT INTO countries (id, code, name) VALUES ($1, $2, $3)", country.ID, country.Code, country.Name)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := c.db.ExecContext(ctx, "UPDATE countries SET code = $2, name = $3 WHERE id = $1", country.ID, country.Code, country.Name)
		if err != nil {
			return nil, err
		}
	}
	return country, nil
}
