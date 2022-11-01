package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
	"go.uber.org/zap"
)

func HandleIndividualDelete(repo db.IndividualRepo) http.Handler {

	const (
		pathParamIndividualID = "individual_id"
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			ctx = r.Context()
			err error
			l   = logging.NewLogger(ctx)
		)

		individual, err := repo.GetByID(ctx, mux.Vars(r)[pathParamIndividualID])
		if err != nil {
			l.Error("failed to get individual", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := repo.SoftDelete(ctx, individual.ID); err != nil {
			l.Error("failed to delete individual", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/countries/%s/individuals", individual.CountryID), http.StatusFound)

	})
}
