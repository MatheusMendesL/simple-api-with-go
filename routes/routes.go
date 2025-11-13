package api

import (
	"_039_projeto3/middle"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// define the routes to the project and calls the handle to check
func ControlRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middle.Cors)

	r.Route("/api", func(r chi.Router) {
		UserRoutes(r)
	})

	return r
}
