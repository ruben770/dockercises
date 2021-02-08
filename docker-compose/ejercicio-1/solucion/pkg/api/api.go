package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func SetRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/people", func(r chi.Router) {
		r.Get("/", getPeople)
		r.Get("/{personId}", getPerson)
	})
	return r
}
