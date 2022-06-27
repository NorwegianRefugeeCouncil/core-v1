package handlers

import (
	"html/template"
	"net/http"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/logging"
	"go.uber.org/zap"

	"github.com/nrc-no/notcore/internal/db"
)

func HandleUsers(templates map[string]*template.Template, repo db.UserRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		const (
			templateName   = "users.gohtml"
			viewParamUsers = "Users"
		)

		var (
			ctx   = r.Context()
			l     = logging.NewLogger(ctx)
			users []*api.User
			err   error
			opts  api.GetAllUsersOptions
		)

		render := func() {
			renderView(templates, templateName, w, r, map[string]interface{}{
				viewParamUsers: users,
			})
		}

		if err := r.ParseForm(); err != nil {
			l.Error("failed to parse form", zap.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		opts, err = api.ParseGetAllUserOptions(r.Form)
		if err != nil {
			l.Error("failed to parse options", zap.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		users, err = repo.GetAll(ctx, opts)
		if err != nil {
			l.Error("failed to get users", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render()

	})
}
