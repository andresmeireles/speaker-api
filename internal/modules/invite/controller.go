package invite

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/andresmeireles/speaker/internal/logger"
	"github.com/andresmeireles/speaker/internal/modules/config"
	"github.com/andresmeireles/speaker/internal/modules/person"
	"github.com/andresmeireles/speaker/internal/web"
	"github.com/go-chi/chi/v5"
)

func Create(w http.ResponseWriter, r *http.Request) {
	invite, err := web.DecodePostBody[InvitePost](r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error("error on create invite controller, cannot decode", err)
		w.Write([]byte(err.Error()))

		return
	}

	_, err = CreateInvite(
		InviteRepository{},
		person.PersonRepository{},
		invite,
	)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error("error on create invite controller, cannot create", err)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Invite successfully created"))
}

func GetAllInvites(w http.ResponseWriter, r *http.Request) {
	repo := InviteRepository{}
	invites, err := repo.GetAll()

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

func Update(inviteId int, w http.ResponseWriter, r *http.Request) {
	invite, err := web.DecodePostBody[InvitePost](r.Body)
	if err != nil {
		logger.Error("error cannot decode", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	if err = UpdateInvite(
		InviteRepository{},
		person.PersonRepository{},
		invite,
		inviteId,
	); err != nil {
		logger.Error("error cannot update", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Invite successfully updated"))
}

func SendInvite(w http.ResponseWriter, r *http.Request) {
	inviteId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(inviteId)

	if err != nil {
		logger.Error("error on send invite controller, cannot parse id", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	inviteText, err := ParseInviteWithTemplate(
		InviteRepository{},
		config.ConfigRepository{},
		id,
	)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	remeberMessage, err := ParseRememberMessage(
		InviteRepository{},
		config.ConfigRepository{},
		id,
	)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	response, err := json.Marshal(map[string]string{
		"invite":   inviteText,
		"remember": remeberMessage,
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

func DeleteInvite(w http.ResponseWriter, r *http.Request) {
	inviteIdParam := chi.URLParam(r, "id")
	inviteId, err := strconv.Atoi(inviteIdParam)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error("error on delete invite controller, bad formatted url", err)
		w.Write([]byte("bad formatted url"))

		return
	}

	repository := InviteRepository{}
	err = RemoveInvite(inviteId, repository)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error("error on delete invite controller, when remove invite", err)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Invite successfully deleted"))
}

func Accepted(w http.ResponseWriter, r *http.Request) {
	inviteIdParam := chi.URLParam(r, "id")
	inviteId, err := strconv.Atoi(inviteIdParam)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error("error on accepted invite controller, error on decode", err)
		w.Write([]byte("bad formatted url"))

		return
	}

	repository := InviteRepository{}
	err = AcceptInvite(inviteId, repository)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad formatted url"))

		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Invite successfully accepted"))
}

func Remember(w http.ResponseWriter, r *http.Request) {
	inviteIdParam := chi.URLParam(r, "id")
	inviteId, err := strconv.Atoi(inviteIdParam)

	if err != nil {
		slog.Error("error on accepted invite controller, error on decode", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad formatted url"))

		return
	}

	repository := InviteRepository{}
	err = RememberInvite(inviteId, repository)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad formatted url"))

		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Invite successfully remembered"))
}
