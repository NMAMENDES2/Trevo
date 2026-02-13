package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/NMAMENDES2/Trevo/api/models"
	"github.com/NMAMENDES2/Trevo/db"
)

type UserHandler struct {
	db *db.Database
}

func NewUserHandler(database *db.Database) *UserHandler {
	return &UserHandler{db: database}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.Pool.Query(r.Context(), "SELECT id, name, email FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
