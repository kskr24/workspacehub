package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/heath", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("OK"))
	})
	return r
}
