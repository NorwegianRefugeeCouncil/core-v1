package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/nrc-no/notcore/internal/api/enumTypes"
	"github.com/nrc-no/notcore/internal/logging"
	"go.uber.org/zap"
)

type CountryRepo interface {
	GetAll(ctx context.Context) ([]*enumTypes.Country, error)
	GetByID(ctx context.Context, id string) (*enumTypes.Country, error)
	Put(ctx context.Context, country *enumTypes.Country) (*enumTypes.Country, error)
}

type countryRepo struct {
	db *sqlx.DB
}

func NewCountryRepo(db *sqlx.DB) CountryRepo {
	return &countryRepo{db: db}
}

func (c countryRepo) logger(ctx context.Context) *zap.Logger {
	return logging.NewLogger(ctx)
}

func (c countryRepo) GetAll(ctx context.Context) ([]*enumTypes.Country, error) {
	l := c.logger(ctx)
	l.Debug("getting all countries")

	const query = "SELECT * FROM countries ORDER BY name"

	auditDuration := logDuration(ctx, "get all countries")
	defer auditDuration()

	var countries []*enumTypes.Country
	if err := c.db.SelectContext(ctx, &countries, query); err != nil {
		l.Error("failed to get countries", zap.Error(err))
		return nil, err
	}
	return countries, nil
}

func (c countryRepo) GetByID(ctx context.Context, id string) (*enumTypes.Country, error) {
	l := c.logger(ctx).With(zap.String("country_id", id))
	l.Debug("getting country by id", zap.String("id", id))

	const query = "SELECT * FROM countries WHERE id = $1"
	var args = []interface{}{id}

	auditDuration := logDuration(ctx, "get country by id")
	defer auditDuration()

	var country enumTypes.Country
	if err := c.db.GetContext(ctx, &country, query, args...); err != nil {
		l.Error("failed to get country by id", zap.Error(err))
		return nil, err
	}
	return &country, nil
}

func (c countryRepo) Put(ctx context.Context, country *enumTypes.Country) (*enumTypes.Country, error) {
	if country.ID == "" {
		return c.createCountry(ctx, country)
	} else {
		return c.updateCountry(ctx, country)
	}
}

func (c countryRepo) updateCountry(ctx context.Context, country *enumTypes.Country) (*enumTypes.Country, error) {
	l := c.logger(ctx)
	l.Debug("updating country")

	const query = "UPDATE countries SET code = $2, name = $3, nrc_organisations = $4 WHERE id = $1"
	var args = []interface{}{
		country.ID,
		country.Code,
		country.Name,
		country.NrcOrganisations,
	}

	auditDuration := logDuration(ctx, "update country")
	defer auditDuration()

	if _, err := c.db.ExecContext(ctx, query, args...); err != nil {
		l.Error("failed to update country", zap.Error(err))
		return nil, err
	}

	return country, nil
}

func (c countryRepo) createCountry(ctx context.Context, country *enumTypes.Country) (*enumTypes.Country, error) {
	l := c.logger(ctx)
	l.Debug("creating new country")
	country.ID = uuid.New().String()

	const query = `INSERT INTO countries (id, code, name, nrc_organisations) VALUES ($1, $2, $3, $4)`

	var args = []interface{}{
		country.ID,
		country.Code,
		country.Name,
		country.NrcOrganisations,
	}

	auditDuration := logDuration(ctx, "create country")
	defer auditDuration()

	if _, err := c.db.ExecContext(ctx, query, args...); err != nil {
		l.Error("failed to create country", zap.Error(err))
		return nil, err
	}

	return country, nil
}
