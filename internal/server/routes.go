package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nrc-no/notcore/internal/db"
	"github.com/nrc-no/notcore/internal/handlers"
	"github.com/nrc-no/notcore/web"
)

func buildRouter(individualRepo db.IndividualRepo, tpl templates) *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/static").Handler(http.FileServer(http.FS(web.Static)))
	r.Path("/").Handler(handlers.HandleIndex(tpl))
	r.Path("/individuals").Handler(handlers.ListHandler(tpl, individualRepo))
	r.Path("/bulk/individuals").Handler(handlers.UploadHandler(individualRepo))
	r.Path("/individuals/{individual_id}").Handler(handlers.HandleIndividual(tpl, individualRepo))
	r.Path("/download").Handler(handlers.HandleDownload(individualRepo))
	return r
}
