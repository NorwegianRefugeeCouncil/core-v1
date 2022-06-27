package handlers

import (
	"github.com/nrc-no/notcore/internal/api"
	"html/template"
	"net/http"

	"github.com/nrc-no/notcore/internal/db"
)

func HandleUsers(templates map[string]*template.Template, repo db.UserRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		opts, err := api.ParseGetAllUserOptions(r.Form)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		users, err := repo.GetAll(ctx, opts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := templates["users.gohtml"].ExecuteTemplate(w, "base", map[string]interface{}{
			"Users": users,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
