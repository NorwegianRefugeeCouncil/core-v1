package server

import (
	"log"
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

	staticRouter := r.PathPrefix("/static").Subrouter()
	staticRouter.HandleFunc("/{file:.*}", web.ServeStatic)

	webRouter := r.PathPrefix("/").Subrouter()
	webRouter.Use(noCache, logHeaders)
	webRouter.Path("/").Handler(handlers.HandleIndex(tpl))
	webRouter.Path("/individuals").Handler(handlers.ListHandler(tpl, individualRepo, countryRepo))
	webRouter.Path("/individuals/upload").Handler(handlers.UploadHandler(individualRepo))
	webRouter.Path("/individuals/download").Handler(handlers.HandleDownload(individualRepo))
	webRouter.Path("/individuals/{individual_id}").Handler(handlers.HandleIndividual(tpl, individualRepo, countryRepo))
	webRouter.Path("/countries").Handler(handlers.HandleCountries(tpl, countryRepo))
	webRouter.Path("/countries/{country_id}").Handler(handlers.HandleCountry(tpl, countryRepo))
	return r
}

func logHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.URL.Path, r.Proto)
		for k, v := range r.Header {
			log.Printf("%s: %s", k, v)
		}
		h.ServeHTTP(w, r)
	})
}

func noCache(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "s-maxage=0 no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		h.ServeHTTP(w, r)
	})
}
