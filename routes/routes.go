package api

import (
	"_046_project/middle"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
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
