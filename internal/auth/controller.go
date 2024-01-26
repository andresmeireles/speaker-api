package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andresmeireles/speaker/internal/codesender"
	"github.com/andresmeireles/speaker/internal/tools/env"
	"github.com/andresmeireles/speaker/internal/tools/responses"
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

func (c AuthController) ReceiveEmail(w http.ResponseWriter, r *http.Request) {
	email, err := web.DecodePostBody[EmailForm](r.Body)
	if err != nil {
		responses.DecodeError(w, err)

		return
	}

	hasEmail := c.actions.HasEmail(email.Email)
	if !hasEmail {
		responses.BadResponse(w, fmt.Errorf("email %s not found", email.Email))

		return
	}

	err = c.actions.SendCode(email.Email)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	responses.Accepted(w, []byte("Code sent"))
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

	responses.Accepted(w, responseJson)
}

func (c AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("user_id").(int)
	if !ok {
		responses.Unauthorized(w, fmt.Errorf("invalid token"))

		return
	}

	err := c.actions.Logout(userId)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	responses.Ok(w, []byte("Logged out"))
}

func (c AuthController) DevAuth(w http.ResponseWriter, r *http.Request) {
	isDev := env.IsDev()
	if !isDev {
		responses.BadResponse(w, fmt.Errorf("not in dev mode"))

		return
	}

	auth, err := c.actions.CreateJWT(user.User{Id: 1, Email: "g6oRz@example.com"}, false)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	responses.Created(w, []byte(auth.Hash))
}
