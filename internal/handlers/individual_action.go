package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

func HandleIndividualAction(repo db.IndividualRepo, action db.IndividualAction) http.Handler {

	const (
		pathParamIndividualID = "individual_id"
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			ctx = r.Context()
			err error
			l   = logging.NewLogger(ctx)
		)

		countryID, err := utils.GetSelectedCountryID(ctx)
		if err != nil {
			l.Error("failed to get selected country", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		individual, err := repo.GetByID(ctx, mux.Vars(r)[pathParamIndividualID])
		if err != nil {
			l.Error("failed to get individual", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if individual.CountryID != countryID {
			l.Warn("user trying to delete individual with the wrong country id", zap.String("individual_id", individual.ID))
			http.Error(w, fmt.Sprintf("individual not found: %v", individual.ID), http.StatusNotFound)
			return
		}

		if err := repo.PerformAction(ctx, individual.ID, action); err != nil {
			l.Error("failed to delete individual", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if action == db.DeleteAction {
			http.Redirect(w, r, fmt.Sprintf("/countries/%s/participants", individual.CountryID), http.StatusFound)
		} else {
			http.Redirect(w, r, fmt.Sprintf("/countries/%s/participants/%s", individual.CountryID, individual.ID), http.StatusFound)
		}

	})
}
