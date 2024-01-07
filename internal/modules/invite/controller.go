package invite

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/andresmeireles/speaker/internal/logger"
	"github.com/andresmeireles/speaker/internal/modules/config"
	"github.com/andresmeireles/speaker/internal/modules/person"
	web "github.com/andresmeireles/speaker/internal/web/decoder"
)

type InviteController struct {
	inviteRepository InviteRepository
	personRepository person.PersonRepository
	configRepository config.ConfigRepository
	action           Actions
}

func (i InviteController) Create(w http.ResponseWriter, r *http.Request) {
	invite, err := web.DecodePostBody[InvitePost](r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error("error on create invite controller, cannot decode", err)
		w.Write([]byte(err.Error()))

		return
	}

	_, err = i.action.CreateInvite(invite)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Invite successfully created"))
}

func (i InviteController) GetAllInvites(w http.ResponseWriter, r *http.Request) {
	invites, err := i.inviteRepository.GetAllOrdered("date", true)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	response, err := json.Marshal(invites)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (i InviteController) Update(inviteId int, w http.ResponseWriter, r *http.Request) {
	invite, err := web.DecodePostBody[InvitePost](r.Body)
	if err != nil {
		slog.Error("error cannot decode", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	if err = UpdateInvite(
		i.inviteRepository,
		i.personRepository,
		invite,
		inviteId,
	); err != nil {
		slog.Error("error cannot update", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Invite successfully updated"))
}

func (i InviteController) SendInvite(inviteId int, w http.ResponseWriter, r *http.Request) {
	inviteText, err := i.action.ParseInviteWithTemplate(inviteId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	rememberMessage, err := ParseRememberMessage(
		i.inviteRepository,
		i.configRepository,
		inviteId,
	)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	response, err := json.Marshal(map[string]string{
		"invite":   inviteText,
		"remember": rememberMessage,
	})
	if err != nil {
		logger.Error("error on send invite controller, cannot parse to json", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (i InviteController) DeleteInvite(inviteId int, w http.ResponseWriter, r *http.Request) {
	repository := i.inviteRepository
	err := RemoveInvite(inviteId, repository)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error("error on delete invite controller, when remove invite", err)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Invite successfully deleted"))
}

func (i InviteController) Accepted(inviteId int, w http.ResponseWriter, r *http.Request) {
	err := i.action.acceptInvite(inviteId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad formatted url"))

		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Invite successfully accepted"))
}

func (i InviteController) Remember(inviteId int, w http.ResponseWriter, r *http.Request) {
	err := i.action.rememberInvite(inviteId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad formatted url"))

		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Invite successfully remembered"))
}
