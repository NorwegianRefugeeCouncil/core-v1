package server

import (
	"github.com/nrc-no/notcore/cmd/devinit"
	"github.com/nrc-no/notcore/internal/clients/zanzibar"
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
	permissionRepo db.PermissionRepo,
	tpl templates,
	config devinit.Config,
) *mux.Router {

	r := mux.NewRouter()
	r.Use(requestIdMiddleware)

	staticRouter := r.PathPrefix("/static").Subrouter()
	staticRouter.HandleFunc("/{file:.*}", web.ServeStatic)

	zanzibarClient := zanzibar.NewZanzibarClient(config)

	webRouter := r.PathPrefix("/").Subrouter()
	webRouter.Use(
		noCache,
		logMiddleware,
		authMiddleware(userRepo, zanzibarClient),
		permissionMiddleware(permissionRepo, zanzibarClient),
		firstUserGlobalAdmin(permissionRepo, zanzibarClient),
	)
	webRouter.Path("/").Handler(handlers.HandleIndex(tpl, zanzibarClient))
	webRouter.Path("/individuals").Handler(handlers.ListHandler(tpl, zanzibarClient, individualRepo, countryRepo))
	webRouter.Path("/individuals/upload").Handler(handlers.UploadHandler(zanzibarClient, individualRepo))
	webRouter.Path("/individuals/download").Handler(handlers.HandleDownload(zanzibarClient, individualRepo, countryRepo))
	webRouter.Path("/individuals/{individual_id}").Handler(handlers.HandleIndividual(tpl, zanzibarClient, individualRepo, countryRepo))
	webRouter.Path("/countries").Handler(handlers.HandleCountries(tpl, zanzibarClient, countryRepo))
	webRouter.Path("/countries/{country_id}").Handler(handlers.HandleCountry(tpl, zanzibarClient, countryRepo))
	webRouter.Path("/users").Handler(handlers.HandleUsers(tpl, zanzibarClient, userRepo))
	webRouter.Path("/users/{user_id}").Handler(handlers.HandleUser(tpl, zanzibarClient, countryRepo, userRepo, permissionRepo))
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
