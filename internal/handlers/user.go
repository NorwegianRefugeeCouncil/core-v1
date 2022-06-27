package handlers

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"

	"github.com/nrc-no/notcore/internal/db"
)

func HandleUser(templates map[string]*template.Template, repo db.UserRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userID := mux.Vars(r)["user_id"]

		user, err := repo.GetByID(ctx, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := templates["user.gohtml"].ExecuteTemplate(w, "base", map[string]interface{}{
			"User": user,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
