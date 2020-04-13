package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func NewHttpRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RedirectSlashes)

	router.Route("/wendy", func(router chi.Router) {
		router.Get("/areyouok", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("I'm wendy i'm ok"))
		})
	})

	return router
}
