package handlers

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func HandleSelectCountry() http.Handler {

	const (
		pathParamCountryID = "country_id"
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			ctx       = r.Context()
			l         = logging.NewLogger(ctx)
			err       error
			countryID string
		)

		countryID = mux.Vars(r)[pathParamCountryID]

		authCtx, err := utils.GetAuthContext(ctx)
		if err != nil {
			l.Error("failed to get auth context", zap.Error(err))
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		allowedCountryIDs := authCtx.GetCountryIDsWithReadWritePermissions()
		if !allowedCountryIDs.Contains(countryID) {
			l.Warn("user is not allowed to select country", zap.Error(err))
			http.Error(w, "You are not allowed to select this country", http.StatusForbidden)
			return
		}

		setSelectedCountryCookie(w, countryID)

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

	})
}

func setSelectedCountryCookie(w http.ResponseWriter, countryID string) {
	http.SetCookie(w, &http.Cookie{
		Name:    "core_selectedCountryID",
		Value:   countryID,
		Expires: time.Now().Add(24 * 30 * 12 * time.Hour),
		Path:    "/",
	})
}
