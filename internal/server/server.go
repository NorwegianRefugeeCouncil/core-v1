package server

import (
	"context"
	"github.com/nrc-no/notcore/cmd/devinit"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Options struct {
	Address        string
	DatabaseDriver string
	DatabaseDSN    string
	LogoutURL      string
}

func (o Options) New() (*Server, error) {
	sqlDb, err := sqlx.Connect(o.DatabaseDriver, o.DatabaseDSN)
	if err != nil {
		return nil, err
	}
	sqlDb.SetMaxIdleConns(5)
	sqlDb.SetMaxOpenConns(10)
	sqlDb.SetConnMaxLifetime(time.Minute * 5)

	if err := db.Migrate(context.Background(), sqlDb); err != nil {
		return nil, err
	}

	individualRepo := db.NewIndividualRepo(sqlDb)
	countryRepo := db.NewCountryRepo(sqlDb)
	userRepo := db.NewUserRepo(sqlDb)
	permissionRepo := db.NewPermissionRepo(sqlDb)

	s := &Server{
		address: o.Address,
	}

	tpl, err := parseTemplates(o.LogoutURL)
	if err != nil {
		return nil, err
	}

	var config devinit.Config
	err = config.MakeConfig()
	if err != nil {
		return nil, err
	}
	
	s.router = buildRouter(individualRepo, countryRepo, userRepo, permissionRepo, tpl, config)

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
