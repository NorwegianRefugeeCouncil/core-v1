package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/logging"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type UserCountry struct {
	UserID    string `db:"user_id"`
	CountryID string `db:"country_id"`
	Read      bool   `db:"read"`
	Write     bool   `db:"write"`
	Admin     bool   `db:"admin"`
}

type UserPermissions struct {
	UserID        string `db:"user_id"`
	IsGlobalAdmin bool   `db:"is_global_admin"`
}

type PermissionRepo interface {
	GetPermissionsForUser(ctx context.Context, userID string) (*api.UserPermissions, error)
	SavePermissionsForUser(ctx context.Context, userPermissions *api.UserPermissions) error
	HasAnyGlobalAdmin(ctx context.Context) (bool, error)
}

type permissionRepo struct {
	db *sqlx.DB
}

func NewPermissionRepo(db *sqlx.DB) PermissionRepo {
	return &permissionRepo{db: db}
}

func (u permissionRepo) HasAnyGlobalAdmin(ctx context.Context) (bool, error) {
	l := logging.NewLogger(ctx)
	l.Debug("checking if system has any global admin")

	var count int
	if err := u.db.GetContext(ctx, &count, "select count(*) from user_permissions where is_global_admin = true"); err != nil {
		l.Error("failed to check if user has any global admin", zap.Error(err))
		return false, err
	}
	return count > 0, nil
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

		findUserPermissionsQry := "select * from user_permissions where user_id = $1"
		findUserPermissionsArgs := []interface{}{userPermissions.UserID}
		var userPermissionsResult UserPermissions

		err := tx.GetContext(ctx, &userPermissionsResult, findUserPermissionsQry, findUserPermissionsArgs...)

		if err == nil {
			// user permissions already exist, update them
			const updateQry = "update user_permissions set is_global_admin = $1 where user_id = $2"
			updateArgs := []interface{}{userPermissions.IsGlobalAdmin, userPermissions.UserID}
			if _, err := tx.ExecContext(ctx, updateQry, updateArgs...); err != nil {
				l.Error("failed to update user permissions", zap.Error(err))
				return nil, err
			}
		} else {
			if errors.Is(err, sql.ErrNoRows) {
				// user permissions does not exist, create it
				insertUserPermissionsQry := "insert into user_permissions (user_id, is_global_admin) values ($1, $2)"
				insertUserPermissionsArgs := []interface{}{userPermissions.UserID, userPermissions.IsGlobalAdmin}
				if _, err := tx.ExecContext(ctx, insertUserPermissionsQry, insertUserPermissionsArgs...); err != nil {
					l.Error("failed to insert user permissions", zap.Error(err))
					return nil, err
				}
			} else {
				l.Error("failed to find user permissions", zap.Error(err))
				return nil, err
			}
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

	var userCountries []*UserCountry
	var userPermissions UserPermissions

	errGrp, gCtx := errgroup.WithContext(ctx)
	errGrp.Go(func() error {
		const userCountriesQry = "select * from user_countries where user_id = $1"
		var userCountriesArgs = []interface{}{userID}
		if err := u.db.SelectContext(gCtx, &userCountries, userCountriesQry, userCountriesArgs...); err != nil {
			l.Error("failed to get countries for user", zap.Error(err))
			return err
		}
		return nil
	})
	errGrp.Go(func() error {
		const userPermissionsQry = "select * from user_permissions where user_id = $1"
		var userPermissionsArgs = []interface{}{userID}
		if err := u.db.GetContext(gCtx, &userPermissions, userPermissionsQry, userPermissionsArgs...); err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				l.Error("failed to get permissions for user", zap.Error(err))
				return err
			}
			insertQry := "insert into user_permissions (user_id, is_global_admin) values ($1, $2)"
			insertArgs := []interface{}{userID, false}
			if _, err := u.db.ExecContext(gCtx, insertQry, insertArgs...); err != nil {
				l.Error("failed to insert permissions for user", zap.Error(err))
				return err
			}
		}
		return nil
	})
	if err := errGrp.Wait(); err != nil {
		l.Error("failed to get user permissions", zap.Error(err))
		return nil, err
	}

	var ret = api.UserPermissions{
		IsGlobalAdmin:      userPermissions.IsGlobalAdmin,
		UserID:             userID,
		CountryPermissions: api.CountryPermissions{},
	}
	for _, userCountry := range userCountries {
		ret.CountryPermissions[userCountry.CountryID] = api.CountryPermission{
			CountryID: userCountry.CountryID,
			Read:      userCountry.Read,
			Write:     userCountry.Write,
			Admin:     userCountry.Admin,
		}
	}

	return &ret, nil
}
