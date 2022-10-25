package server

import (
	"net/http"

	gorillahandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/auth"
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

	countriesRouter := webRouter.PathPrefix("/countries").Subrouter()
	countriesRouter.Path("").Handler(handlers.HandleCountries(tpl))

	countryRouter := countriesRouter.PathPrefix("/{country_id}").Subrouter()
	countryRouter.Path("").Handler(withMiddleware(
		handlers.HandleCountry(tpl, countryRepo),
		hasGlobalAdminPermissionMiddleware(),
	))

	individualsRouter := countryRouter.PathPrefix("/individuals").Subrouter()
	individualsRouter.Path("").Methods(http.MethodGet).Handler(withMiddleware(
		handlers.HandleIndividuals(renderer, individualRepo),
		ensureSelectedCountryMiddleware(),
		hasCountryPermissionMiddleware(auth.PermissionRead),
	))
	individualsRouter.Path("/upload").Methods(http.MethodPost).Handler(withMiddleware(
		handlers.UploadHandler(individualRepo),
		ensureSelectedCountryMiddleware(),
		hasCountryPermissionMiddleware(auth.PermissionWrite),
	))
	individualsRouter.Path("/download").Methods(http.MethodGet).Handler(withMiddleware(
		handlers.HandleDownload(individualRepo),
		ensureSelectedCountryMiddleware(),
		hasCountryPermissionMiddleware(auth.PermissionRead),
	))
	individualsRouter.Path("/{individual_id}").Methods(http.MethodGet).Handler(withMiddleware(
		handlers.HandleIndividual(tpl, individualRepo),
		ensureSelectedCountryMiddleware(),
		hasCountryPermissionMiddleware(auth.PermissionRead),
	))
	individualsRouter.Path("/{individual_id}").Methods(http.MethodPost).Handler(withMiddleware(
		handlers.HandleIndividual(tpl, individualRepo),
		ensureSelectedCountryMiddleware(),
		hasCountryPermissionMiddleware(auth.PermissionWrite),
	))

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
