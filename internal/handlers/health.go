package handlers

import (
	"net/http"

	"github.com/nrc-no/notcore/internal/db"
)

func HandleHealth(healthzRepo db.HealthzRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := healthzRepo.Check(r.Context())

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
