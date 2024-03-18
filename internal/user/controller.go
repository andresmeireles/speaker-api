package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andresmeireles/speaker/internal/tools/responses"
)

type UserController struct {
	repository Repository
}

func NewController(repository Repository) UserController {
	return UserController{repository}
}

func (c UserController) Me(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("user_id").(int)
	if !ok {
		responses.BadResponse(w, fmt.Errorf("invalid user id %s", userId))

		return
	}

	user, err := c.repository.GetById(userId)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	responses.Ok(w, jsonUser)
}
