package handlers

import (
	"net/http"

	"github.com/NMAMENDES2/Trevo/api/models"
	"github.com/NMAMENDES2/Trevo/db"
	"github.com/NMAMENDES2/Trevo/pkg/response"
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

	response.JSON(w, http.StatusOK, users)
}
