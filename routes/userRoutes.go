package api

import (
	"_046_project/response"

	"github.com/go-chi/chi"
)

func UserRoutes(r chi.Router) {

	r.Route("/user", func(r chi.Router) {
		r.Get("/", response.GetAllUsers)
		/* r.Get("/{id}", response.GetByID)
		r.Post("/", response.AddUser)
		r.Put("/{id}", response.EditUser)
		r.Delete("/{id}", response.DeleteUser)
		r.Get("/search/{name}", response.SearchUser) */
	})
}
