package auth

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/andresmeireles/speaker/internal/codesender"
	"github.com/andresmeireles/speaker/internal/tools/responses"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
	"github.com/andresmeireles/speaker/internal/user"
	web "github.com/andresmeireles/speaker/internal/web/decoder"
)

type AuthController struct {
	actions           Actions
	codesenderActions codesender.Actions
	userRepository    user.Repository
}

func NewController(action Actions, codeSenderActions codesender.Actions, userRepo user.Repository) AuthController {
	return AuthController{
		actions:           action,
		codesenderActions: codeSenderActions,
		userRepository:    userRepo,
	}
}

func (c AuthController) New(s servicelocator.ServiceLocator) any {
	c.actions = servicelocator.Get[Actions](s)
	c.codesenderActions = servicelocator.Get[codesender.Actions](s)
	c.userRepository = servicelocator.Get[user.Repository](s)

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
		responses.DecodeError(w, err)

		return
	}

	err = c.codesenderActions.VerifyCode(form.Email, form.Code)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	user, err := c.userRepository.GetByEmail(form.Email)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	jwt, err := c.actions.CreateJWT(user, form.Remember)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	response := map[string]string{
		"token":   jwt.Hash,
		"message": "Token Created",
	}
	responseJson, err := json.Marshal(response)

	if err != nil {
		responses.BadResponse(w, err)

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
