package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ramosjanoah/eidolmou/wendy/service"
	"net/http"
)

var (
	AppURL       = "https://white-wendy.herokuapp.com"
	HeartbeatURL = AppURL + "/wendy/areyouok"
)

func NewHttpRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RedirectSlashes)

	router.Route("/wendy", func(router chi.Router) {
		router.Get("/areyouok", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(service.AreYouOKResponseMsg))
		})
	})

	return router
}
