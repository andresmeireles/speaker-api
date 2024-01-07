package user

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type UserController struct {
	repository UserRepository
}

func (c UserController) Me(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(int)
	user, err := c.repository.GetById(userId)

	if err != nil {
		slog.Error("No user id on request", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		slog.Error("Failed to marshal user", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Write(jsonUser)
}
