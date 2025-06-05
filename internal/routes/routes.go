package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kskr24/workspacehub/internal/handlers"
)

func RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/users/{id}", handlers.GetUserByIDHandler)
	r.Put("/users/{id}", handlers.UpdateUser)
	return r
}
