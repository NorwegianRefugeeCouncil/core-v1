package server

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/server/middleware"
	"go.uber.org/zap"
)

func (o Options) New(ctx context.Context) (*Server, error) {

	l := logging.NewLogger(ctx)

	if err := o.validate(); err != nil {
		l.Error("invalid options", zap.Error(err))
		return nil, err
	}
	sqlDb, err := sqlx.ConnectContext(ctx, o.DatabaseDriver, o.DatabaseDSN)
	if err != nil {
		l.Error("failed to connect to db", zap.Error(err))
		return nil, err
	}
	// TODO: make this configurable at some point
	sqlDb.SetMaxIdleConns(5)
	sqlDb.SetMaxOpenConns(10)
	sqlDb.SetConnMaxLifetime(time.Minute * 5)

	if err := db.Migrate(context.Background(), sqlDb); err != nil {
		l.Error("failed to migrate database", zap.Error(err))
		return nil, err
	}

	// create the individual db repository
	individualRepo := db.NewIndividualRepo(sqlDb)

	// create the country db repository
	countryRepo := db.NewCountryRepo(sqlDb)

	s := &Server{address: o.Address}

	// parse html templates
	tpl, err := parseTemplates(o.LogoutURL, o.TokenRefreshURL)
	if err != nil {
		l.Error("failed to parse templates", zap.Error(err))
		return nil, err
	}

	// create the oidc provider
	oidcProvider, err := oidc.NewProvider(ctx, o.OIDCIssuerURL)
	if err != nil {
		l.Error("failed to get oidc provider", zap.Error(err))
		return nil, err
	}

	// create the id token verifier
	oidcVerifier := oidcProvider.Verifier(&oidc.Config{
		ClientID:             o.OAuthClientID,
		SupportedSigningAlgs: []string{oidc.RS256},
	})
	idTokenVerifier := middleware.NewIDTokenVerifier(oidcVerifier)

	// build the router
	s.router = buildRouter(
		individualRepo,
		countryRepo,
		o.JwtGroupGlobalAdmin,
		o.AuthHeaderName,
		o.AuthHeaderFormat,
		o.LoginURL,
		idTokenVerifier,
		tpl,
	)

	return s, nil
}

type Server struct {
	address  string
	listener net.Listener
	router   *mux.Router
}

func (s *Server) Start(ctx context.Context) error {
	l := logging.NewLogger(ctx)
	l.Info("starting server")
	var err error

	s.listener, err = net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	l.Info("listening on " + s.listener.Addr().String())

	go func() {
		<-ctx.Done()
		l.Info("stopping server")
		s.listener.Close()
	}()
	err = http.Serve(s.listener, s.router)
	l.Info("server stopped")
	if err != nil {
		return err
	}
	return nil
}
