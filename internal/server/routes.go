package server

import (
	"net/http"

	gorillahandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/handlers"
	"github.com/nrc-no/notcore/web"
)

func buildRouter(
	individualRepo db.IndividualRepo,
	countryRepo db.CountryRepo,
	globalAdminGroup string,
	authHeaderName string,
	authHeaderFormat string,
	loginURL string,
	idTokenVerifier IDTokenVerifier,
	tpl templates,
) *mux.Router {

	r := mux.NewRouter()
	r.Use(
		gorillahandlers.RecoveryHandler(),
		gorillahandlers.CompressHandler,
		requestIdMiddleware,
	)
	renderer := handlers.NewRenderer(tpl)

	staticRouter := r.PathPrefix("/static").Subrouter()
	staticRouter.HandleFunc("/{file:.*}", web.ServeStatic)

	webRouter := r.PathPrefix("").Subrouter()

	webRouter.Use(
		noCache,
		logMiddleware,
		authMiddleware(authHeaderName, authHeaderFormat, idTokenVerifier, loginURL),
		countriesMiddleware(countryRepo),
		permissionMiddleware(globalAdminGroup),
		selectedCountryMiddleware(),
	)

	webRouter.Path("/api/session").Handler(handlers.HandleSession())

	webRouter.Path("/countries").Handler(handlers.HandleCountries(tpl))
	webRouter.Path("/countries/{country_id}").Handler(handlers.HandleCountry(tpl, countryRepo))

	webRouter.Path("/countries/{country_id}/individuals").Handler(withMiddleware(
		handlers.HandleIndividuals(renderer, individualRepo),
		ensureSelectedCountryMiddleware()),
	)

	webRouter.Path("/countries/{country_id}/individuals/upload").Handler(withMiddleware(
		handlers.UploadHandler(individualRepo),
		ensureSelectedCountryMiddleware()),
	)

	webRouter.Path("/countries/{country_id}/individuals/download").Handler(withMiddleware(
		handlers.HandleDownload(individualRepo),
		ensureSelectedCountryMiddleware()),
	)

	webRouter.Path("/countries/{country_id}/individuals/{individual_id}").Handler(handlers.HandleIndividual(tpl, individualRepo))

	webRouter.PathPrefix("").Handler(handlers.HandleHome(tpl))

	return r
}

func withMiddleware(handler http.Handler, mw ...mux.MiddlewareFunc) http.Handler {
	for i := len(mw) - 1; i >= 0; i-- {
		handler = mw[i](handler)
	}
	return handler
}

func noCache(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "s-maxage=0 no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		h.ServeHTTP(w, r)
	})
}
