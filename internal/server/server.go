package server

import (
	"context"
	"encoding/hex"
	"net"
	"net/http"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/locales"
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

	if err := locales.LoadTranslations(); err != nil {
		l.Error("failed to load translations", zap.Error(err))
		return nil, err
	}

	// create the healthz db repository
	healthzRepo := db.NewHealthzRepo(sqlDb)

	// create the individual db repository
	individualRepo := db.NewIndividualRepo(sqlDb)

	// create the country db repository
	countryRepo := db.NewCountryRepo(sqlDb)

	s := &Server{address: o.Address}

	// parse html templates
	tpl, err := parseTemplates(o.LoginURL, o.TokenRefreshURL, o.TokenRefreshInterval)
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

	hashKey1, err := hex.DecodeString(o.HashKey1)
	if err != nil {
		l.Error("failed to decode hash key 1", zap.Error(err))
		return nil, err
	}

	blockKey1, err := hex.DecodeString(o.BlockKey1)
	if err != nil {
		l.Error("failed to decode block key 1", zap.Error(err))
		return nil, err
	}

	hashKey2, err := hex.DecodeString(o.HashKey2)
	if err != nil {
		l.Error("failed to decode hash key 2", zap.Error(err))
		return nil, err
	}

	blockKey2, err := hex.DecodeString(o.BlockKey2)
	if err != nil {
		l.Error("failed to decode block key 2", zap.Error(err))
		return nil, err
	}

	sessionStore := sessions.NewCookieStore(
		hashKey1,
		blockKey1,
		hashKey2,
		blockKey2,
	)
	sessionStore.MaxAge(60 * 60) // 1 hour
	sessionStore.Options.HttpOnly = true
	sessionStore.Options.Secure = true
	sessionStore.Options.SameSite = http.SameSiteStrictMode

	l.Info("Server Option, EnableBetaFeatures: ", zap.Bool("enableBetaFeatures", o.EnableBetaFeatures))

	// build the router
	s.router = buildRouter(
		healthzRepo,
		individualRepo,
		countryRepo,
		o.JwtGroups,
		o.IdTokenAuthHeaderName,
		o.IdTokenAuthHeaderFormat,
		o.AccessTokenHeaderName,
		o.AccessTokenHeaderFormat,
		o.LoginURL,
		o.EnableBetaFeatures,
		oidcProvider,
		idTokenVerifier,
		sessionStore,
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
