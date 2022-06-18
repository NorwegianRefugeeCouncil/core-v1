package server

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/nrc-no/notcore/internal/db"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Options struct {
	Address        string
	DatabaseDriver string
	DatabaseDSN    string
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

	s := &Server{
		address: o.Address,
	}

	tpl, err := parseTemplates()
	if err != nil {
		return nil, err
	}

	s.router = buildRouter(individualRepo, countryRepo, tpl)

	return s, nil
}

type Server struct {
	address  string
	listener net.Listener
	router   *mux.Router
}

func (s *Server) Start(ctx context.Context) error {
	var err error

	s.listener, err = net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		s.listener.Close()
	}()
	return http.Serve(s.listener, s.router)
}
