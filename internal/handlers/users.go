package handlers

import (
	"github.com/nrc-no/notcore/internal/clients/zanzibar"
	"html/template"
	"net/http"

	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"

	"github.com/nrc-no/notcore/internal/db"
)

func HandleUsers(templates map[string]*template.Template, client zanzibar.Client, userRepo db.UserRepo) http.Handler {

	const (
		templateName   = "users.gohtml"
		viewParamUsers = "Users"
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			ctx   = r.Context()
			l     = logging.NewLogger(ctx)
			users []*api.User
			err   error
			opts  api.GetAllUsersOptions
		)

		render := func() {
			renderView(templates, templateName, w, r, viewParams{
				viewParamUsers: users,
			})
		}

		if !utils.IsGlobalAdmin(ctx) {
			l.Warn("cannot access country page without global admin role")
			http.Error(w, "user is not global admin", http.StatusForbidden)
			return
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

		users, err = userRepo.GetAll(ctx, opts)
		if err != nil {
			l.Error("failed to get users", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render()

	})
}
