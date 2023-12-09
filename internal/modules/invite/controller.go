package invite

import (
	"encoding/json"
	"net/http"

	"github.com/andresmeireles/speaker/internal/modules/config"
	"github.com/andresmeireles/speaker/internal/modules/person"
	"github.com/andresmeireles/speaker/internal/web"
)

func Create(w http.ResponseWriter, r *http.Request) {
	invite, err := web.DecodePostBody[InvitePost](r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	_, err = CreateInvite(
		InviteRepository{},
		person.PersonRepository{},
		invite,
	)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Invite successfully created"))
}

func GetAllInvites(w http.ResponseWriter, r *http.Request) {
	repo := InviteRepository{}
	invites := repo.GetAll()
	response, err := json.Marshal(invites)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Update(inviteId int, w http.ResponseWriter, r *http.Request) {
	invite, err := web.DecodePostBody[InvitePost](r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	_, err = UpdateInvite(
		InviteRepository{},
		person.PersonRepository{},
		invite,
		inviteId,
	)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Invite successfully updated"))
}

func SendInvite(w http.ResponseWriter, r *http.Response) {
	body, err := web.DecodePostBody[InviteSender](r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	inviteText, err := ParseInviteWithTemplate(
		InviteRepository{},
		config.ConfigRepository{},
		body,
	)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(inviteText))
}
