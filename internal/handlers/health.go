package handlers

import (
	"net/http"

	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/logging"
)

func HandleHealth(healthzRepo db.HealthzRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := logging.NewLogger(ctx)

		err := healthzRepo.Check(r.Context())

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		l.Info("health check ok")

		w.WriteHeader(http.StatusOK)
	})
}
