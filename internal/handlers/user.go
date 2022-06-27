package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/logging"
	"go.uber.org/zap"

	"github.com/nrc-no/notcore/internal/db"
)

func HandleUser(templates map[string]*template.Template, repo db.UserRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		const (
			templateName    = "user.gohtml"
			pathParamUserId = "user_id"
			viewParamUser   = "User"
		)

		var (
			ctx    = r.Context()
			err    error
			userID = mux.Vars(r)[pathParamUserId]
			user   *api.User
			l      = logging.NewLogger(ctx)
		)

		render := func() {
			renderView(templates, templateName, w, r, map[string]interface{}{
				viewParamUser: user,
			})
		}

		user, err = repo.GetByID(ctx, userID)
		if err != nil {
			l.Error("failed to get user", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render()

	})
}
