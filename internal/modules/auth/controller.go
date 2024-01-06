package auth

import (
	"log/slog"
	"net/http"

	web "github.com/andresmeireles/speaker/internal/web/decoder"
)

type AuthController struct {
	actions Actions
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
		return
	}

	// SendCode()

	// err := triggerValidationCode(email.Email)
	// if err != nil {
	// 	slog.Error("Failed to decode email", err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("error"))

	// 	return
	// }

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Validation code was send to email"))
}

// func Login(response http.ResponseWriter, request *http.Request) {
// 	body, err := json.NewDecoder().Decode()
// }

// func Logout(response http.ResponseWriter, request *http.Request) {

// }
