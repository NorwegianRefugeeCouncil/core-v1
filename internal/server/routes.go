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
	userRepo db.UserRepo,
	tpl templates,
) *mux.Router {

	r := mux.NewRouter()
	r.Use(requestIdMiddleware)

	staticRouter := r.PathPrefix("/static").Subrouter()
	staticRouter.HandleFunc("/{file:.*}", web.ServeStatic)

	webRouter := r.PathPrefix("/").Subrouter()
	webRouter.Use(noCache, logMiddleware, authMiddleware(userRepo))
	webRouter.Path("/").Handler(handlers.HandleIndex(tpl))
	webRouter.Path("/individuals").Handler(handlers.ListHandler(tpl, individualRepo, countryRepo))
	webRouter.Path("/individuals/upload").Handler(handlers.UploadHandler(individualRepo))
	webRouter.Path("/individuals/download").Handler(handlers.HandleDownload(individualRepo))
	webRouter.Path("/individuals/{individual_id}").Handler(handlers.HandleIndividual(tpl, individualRepo, countryRepo))
	webRouter.Path("/countries").Handler(handlers.HandleCountries(tpl, countryRepo))
	webRouter.Path("/countries/{country_id}").Handler(handlers.HandleCountry(tpl, countryRepo))
	webRouter.Path("/users").Handler(handlers.HandleUsers(tpl, userRepo))
	webRouter.Path("/users/{user_id}").Handler(handlers.HandleUser(tpl, userRepo))
	return r
}

func noCache(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "s-maxage=0 no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		h.ServeHTTP(w, r)
	})
}
