package auth

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/andresmeireles/speaker/internal/codesender"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
	"github.com/andresmeireles/speaker/internal/user"
	web "github.com/andresmeireles/speaker/internal/web/decoder"
)

type AuthController struct {
	actions           Actions
	codesenderActions codesender.Actions
	userRepository    user.UserRepository
}

func (c AuthController) New(s servicelocator.ServiceLocator) any {
	c.actions = servicelocator.Get[Actions](s)
	c.codesenderActions = servicelocator.Get[codesender.Actions](s)
	c.userRepository = servicelocator.Get[user.UserRepository](s)

	return AuthController{
		actions:           c.actions,
		codesenderActions: c.codesenderActions,
		userRepository:    c.userRepository,
	}
}

func (c AuthController) ReceiveEmail(w http.ResponseWriter, r *http.Request) {
	email, err := web.DecodePostBody[EmailForm](r.Body)
	if err != nil {
		slog.Error("Failed to decode email", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error"))

		return
	}

	hasEmail := c.actions.HasEmail(email.Email)
	if !hasEmail {
		slog.Error("Email not found", "email", email.Email)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("email not found"))

		return
	}

	err = c.actions.SendCode(email.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Validation code was send to email"))
}

func (c AuthController) ReceiveCode(w http.ResponseWriter, r *http.Request) {
	form, err := web.DecodePostBody[CodeForm](r.Body)
	if err != nil {
		slog.Error("Failed to decode code", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error"))

		return
	}

	err = c.codesenderActions.VerifyCode(form.Email, form.Code)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	user, err := c.userRepository.GetByEmail(form.Email)
	if err != nil {
		slog.Error("Failed to get user", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	jwt, err := c.actions.CreateJWT(user)
	if err != nil {
		slog.Error("Failed to create jwt", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	response := map[string]string{
		"token":   jwt.Hash,
		"message": "Token Created",
	}
	responseJson, err := json.Marshal(response)

	if err != nil {
		slog.Error("Failed to marshal response", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(responseJson)
}

func (c AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("user_id").(int)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))

		return
	}

	err := c.actions.Logout(userId)
	if err != nil {
		slog.Error("Failed to logout", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logged out"))
}
