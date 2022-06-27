package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/logging"
	"go.uber.org/zap"
)

type UserCountry struct {
	UserID    string `db:"user_id"`
	CountryID string `db:"country_id"`
	Read      bool   `db:"read"`
	Write     bool   `db:"write"`
	Admin     bool   `db:"admin"`
}

type PermissionRepo interface {
	GetPermissionsForUser(ctx context.Context, userID string) (*api.UserPermissions, error)
	SavePermissionsForUser(ctx context.Context, userPermissions *api.UserPermissions) error
}

type permissionRepo struct {
	db *sqlx.DB
}

func NewPermissionRepo(db *sqlx.DB) PermissionRepo {
	return &permissionRepo{db: db}
}

func (u permissionRepo) SavePermissionsForUser(ctx context.Context, userPermissions *api.UserPermissions) error {
	l := logging.NewLogger(ctx).With(zap.String("user_id", userPermissions.UserID))
	l.Debug("saving user permissions")
	_, err := doInTransaction(ctx, u.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {

		// delete all country permissions for user
		const deleteQry = "delete from user_countries where user_id = $1"
		deleteArgs := []interface{}{userPermissions.UserID}

		if _, err := tx.ExecContext(ctx, deleteQry, deleteArgs...); err != nil {
			l.Error("failed to delete user countries", zap.Error(err))
			return nil, err
		}

		// create new country permissions
		var userCountries = make([]*UserCountry, len(userPermissions.CountryPermissions))
		var i = 0
		for _, countryPermission := range userPermissions.CountryPermissions {
			userCountries[i] = &UserCountry{
				UserID:    userPermissions.UserID,
				CountryID: countryPermission.CountryID,
				Read:      countryPermission.Read,
				Write:     countryPermission.Write,
				Admin:     countryPermission.Admin,
			}
			i++
		}

		// if there are no country permissions, return early
		if len(userCountries) == 0 {
			return nil, nil
		}

		// build insertion query
		var insertArgs []interface{}
		var valueLists []string
		for _, userCountry := range userCountries {
			if !userCountry.Read && !userCountry.Write && !userCountry.Admin {
				continue
			}
			valueLists = append(valueLists, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)",
				len(insertArgs)+1,
				len(insertArgs)+2,
				len(insertArgs)+3,
				len(insertArgs)+4,
				len(insertArgs)+5,
			))
			insertArgs = append(insertArgs, userCountry.UserID, userCountry.CountryID, userCountry.Read, userCountry.Write, userCountry.Admin)
		}

		// if there are no permissions, return early
		if len(valueLists) == 0 {
			return nil, nil
		}

		insertQuery := "insert into user_countries (user_id, country_id, read, write, admin) values " + strings.Join(valueLists, ",")

		if _, err := tx.ExecContext(ctx, insertQuery, insertArgs...); err != nil {
			l.Error("failed to insert user countries", zap.Error(err))
			return nil, err
		}

		return nil, nil

	})

	if err != nil {
		l.Error("failed to save user permissions", zap.Error(err))
		return err
	}

	return nil
}

func (u permissionRepo) GetPermissionsForUser(ctx context.Context, userID string) (*api.UserPermissions, error) {
	l := logging.NewLogger(ctx).With(zap.String("user_id", userID))
	l.Debug("getting user countries")

	const query = "select * from user_countries where user_id = $1"
	var args = []interface{}{userID}

	var result []*UserCountry
	if err := u.db.SelectContext(ctx, &result, query, args...); err != nil {
		l.Error("failed to get countries for user", zap.Error(err))
		return nil, err
	}

	var ret = api.UserPermissions{
		UserID:             userID,
		CountryPermissions: api.CountryPermissions{},
	}
	for _, userCountry := range result {
		ret.CountryPermissions[userCountry.CountryID] = api.CountryPermission{
			CountryID: userCountry.CountryID,
			Read:      userCountry.Read,
			Write:     userCountry.Write,
			Admin:     userCountry.Admin,
		}
	}

	return &ret, nil
}
