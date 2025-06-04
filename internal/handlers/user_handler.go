package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kskr24/workspacehub/internal/services"
)

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		http.Error(w, "Invalid User id", http.StatusBadRequest)
		return
	}

	user, err := services.GetUserByID(id)

	if err != nil {
		http.Error(w, "Error fetching user", http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}
