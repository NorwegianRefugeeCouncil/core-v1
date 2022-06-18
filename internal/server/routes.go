package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/handlers"
	"github.com/nrc-no/notcore/web"
)

func buildRouter(
	individualRepo db.IndividualRepo,
	countryRepo db.CountryRepo,
	tpl templates) *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/static").Handler(http.FileServer(http.FS(web.Static)))
	r.Path("/").Handler(handlers.HandleIndex(tpl))
	r.Path("/individuals").Handler(handlers.ListHandler(tpl, individualRepo, countryRepo))
	r.Path("/individuals/upload").Handler(handlers.UploadHandler(individualRepo))
	r.Path("/individuals/download").Handler(handlers.HandleDownload(individualRepo))
	r.Path("/individuals/{individual_id}").Handler(handlers.HandleIndividual(tpl, individualRepo, countryRepo))
	r.Path("/countries").Handler(handlers.HandleCountries(tpl, countryRepo))
	r.Path("/countries/{country_id}").Handler(handlers.HandleCountry(tpl, countryRepo))
	return r
}
