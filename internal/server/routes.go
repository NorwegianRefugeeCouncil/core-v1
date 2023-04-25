package server

import (
	"github.com/nrc-no/notcore/internal/utils"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	gorillahandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/nrc-no/notcore/internal/auth"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/handlers"
	"github.com/nrc-no/notcore/internal/server/middleware"
	"github.com/nrc-no/notcore/web"
)

func buildRouter(
	individualRepo db.IndividualRepo,
	countryRepo db.CountryRepo,
	jwtGroups utils.JwtGroupOptions,
	idTokenAuthHeaderName string,
	idTokenAuthHeaderFormat string,
	accessTokenHeaderName string,
	accessTokenHeaderFormat string,
	loginURL string,
	provider *oidc.Provider,
	idTokenVerifier middleware.IDTokenVerifier,
	sessionStore *sessions.CookieStore,
	tpl templates,
) *mux.Router {

	r := mux.NewRouter()
	r.Use(
		gorillahandlers.RecoveryHandler(gorillahandlers.PrintRecoveryStack(true)),
		// gorillahandlers.CompressHandler,
		middleware.RequestId,
	)
	renderer := handlers.NewRenderer(tpl)

	staticRouter := r.PathPrefix("/static").Subrouter()
	staticRouter.HandleFunc("/{file:.*}", web.ServeStatic)

	r.Path("/health").Handler(handlers.HandleHealth())

	webRouter := r.PathPrefix("").Subrouter()
	webRouter.Use(
		noCache,
		middleware.RequestLogging,
		middleware.Authentication(idTokenAuthHeaderName, idTokenAuthHeaderFormat, accessTokenHeaderName, accessTokenHeaderFormat, provider, idTokenVerifier, sessionStore, loginURL),
		middleware.PrefetchCountries(countryRepo),
		middleware.ComputePermissions(jwtGroups),
		middleware.SelectedCountry(),
	)

	countriesRouter := webRouter.PathPrefix("/countries").Subrouter()
	countriesRouter.Path("").Handler(handlers.HandleCountries(tpl))

	countryRouter := countriesRouter.PathPrefix("/{country_id}").Subrouter()
	countryRouter.Path("").Handler(withMiddleware(
		handlers.HandleCountry(tpl, countryRepo),
		middleware.HasGlobalAdminPermission(),
	))

	individualsRouter := countryRouter.PathPrefix("/participants").Subrouter()
	individualsRouter.Path("").Methods(http.MethodGet).Handler(withMiddleware(
		handlers.HandleIndividuals(renderer, individualRepo),
		middleware.EnsureSelectedCountry(),
		middleware.HasCountryPermission(auth.PermissionRead),
	))
	individualsRouter.Path("/upload").Methods(http.MethodPost).Handler(withMiddleware(
		handlers.HandleUpload(renderer, individualRepo),
		middleware.EnsureSelectedCountry(),
		middleware.HasCountryPermission(auth.PermissionWrite),
	))
	individualsRouter.Path("/download").Methods(http.MethodGet).Handler(withMiddleware(
		handlers.HandleDownload(individualRepo),
		middleware.EnsureSelectedCountry(),
		middleware.HasCountryPermission(auth.PermissionRead),
	))
	individualsRouter.Path("/delete").Methods(http.MethodPost).Handler(withMiddleware(
		handlers.HandleIndividualsAction(renderer, individualRepo, db.DeleteAction),
		middleware.EnsureSelectedCountry(),
		middleware.HasCountryPermission(auth.PermissionWrite),
	))
	individualsRouter.Path("/deactivate").Methods(http.MethodPost).Handler(withMiddleware(
		handlers.HandleIndividualsAction(renderer, individualRepo, db.DeactivateAction),
		middleware.EnsureSelectedCountry(),
		middleware.HasCountryPermission(auth.PermissionWrite),
	))
	individualsRouter.Path("/activate").Methods(http.MethodPost).Handler(withMiddleware(
		handlers.HandleIndividualsAction(renderer, individualRepo, db.ActivateAction),
		middleware.EnsureSelectedCountry(),
		middleware.HasCountryPermission(auth.PermissionWrite),
	))

	individualRouter := individualsRouter.PathPrefix("/{individual_id}").Subrouter()
	individualRouter.Path("").Methods(http.MethodGet).Handler(withMiddleware(
		handlers.HandleIndividual(tpl, individualRepo),
		middleware.EnsureSelectedCountry(),
		middleware.HasCountryPermission(auth.PermissionRead),
	))
	individualRouter.Path("").Methods(http.MethodPost).Handler(withMiddleware(
		handlers.HandleIndividual(tpl, individualRepo),
		middleware.EnsureSelectedCountry(),
		middleware.HasCountryPermission(auth.PermissionWrite),
	))
	individualRouter.Path("/delete").Methods(http.MethodPost).Handler(withMiddleware(
		handlers.HandleIndividualAction(individualRepo, db.DeleteAction),
		middleware.EnsureSelectedCountry(),
		middleware.HasCountryPermission(auth.PermissionWrite),
	))
	individualRouter.Path("/activate").Methods(http.MethodPost).Handler(withMiddleware(
		handlers.HandleIndividualAction(individualRepo, db.ActivateAction),
		middleware.EnsureSelectedCountry(),
		middleware.HasCountryPermission(auth.PermissionWrite),
	))
	individualRouter.Path("/deactivate").Methods(http.MethodPost).Handler(withMiddleware(
		handlers.HandleIndividualAction(individualRepo, db.DeactivateAction),
		middleware.EnsureSelectedCountry(),
		middleware.HasCountryPermission(auth.PermissionWrite),
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
