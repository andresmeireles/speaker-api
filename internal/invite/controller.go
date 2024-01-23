package invite

import (
	"encoding/json"
	"net/http"

	"github.com/andresmeireles/speaker/internal/config"
	"github.com/andresmeireles/speaker/internal/person"
	"github.com/andresmeireles/speaker/internal/tools/responses"
	web "github.com/andresmeireles/speaker/internal/web/decoder"
)

type InviteController struct {
	inviteRepository InviteRepository
	personRepository person.PersonRepository
	configRepository config.ConfigRepository
	service          InviteService
}

func NewController(
	ir InviteRepository,
	pr person.PersonRepository,
	cr config.ConfigRepository,
	a InviteService,
) InviteController {
	return InviteController{
		inviteRepository: ir,
		personRepository: pr,
		configRepository: cr,
		service:          a,
	}
}

func (i InviteController) Create(w http.ResponseWriter, r *http.Request) {
	invite, err := web.DecodePostBody[InvitePost](r.Body)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	_, err = i.service.CreateInvite(invite)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	responses.Created(w, "Invite successfully created")
}

func (i InviteController) GetInvite(inviteId int, w http.ResponseWriter, r *http.Request) {
	invite, err := i.inviteRepository.GetById(inviteId)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	response, err := json.Marshal(invite)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	responses.Ok(w, string(response))
}

func (i InviteController) GetAllInvites(w http.ResponseWriter, r *http.Request) {
	invites, err := i.inviteRepository.GetAllOrdered("date", true)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	response, err := json.Marshal(invites)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	responses.Ok(w, string(response))
}

func (i InviteController) Update(inviteId int, w http.ResponseWriter, r *http.Request) {
	inviteUpdateData, err := web.DecodePostBody[UpdateInviteData](r.Body)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	if err = i.service.UpdateInvite(
		inviteUpdateData,
		inviteId,
	); err != nil {
		responses.BadResponse(w, err)

		return
	}

	responses.Accepted(w, "Invite successfully updated")
}

func (i InviteController) SendInvite(inviteId int, w http.ResponseWriter, r *http.Request) {
	inviteText, err := i.service.ParseInviteWithTemplate(inviteId)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	rememberMessage, err := i.service.ParseRememberMessage(inviteId)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	response, err := json.Marshal(map[string]string{
		"invite":   inviteText,
		"remember": rememberMessage,
	})
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	responses.Created(w, string(response))
}

func (i InviteController) Accepted(inviteId int, w http.ResponseWriter, r *http.Request) {
	err := i.service.AcceptInvite(inviteId)

	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	responses.Ok(w, "Invite successfully accepted")
}

func (c InviteController) Remember(inviteId int, w http.ResponseWriter, r *http.Request) {
	err := c.service.RememberInvite(inviteId)

	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	responses.Ok(w, "Invite successfully remembered")
}

func (c InviteController) WasDone(inviteId int, w http.ResponseWriter, r *http.Request) {
	wasDone, err := web.DecodePostBody[WasDone](r.Body)
	if err != nil {
		responses.DecodeError(w, err)

		return
	}

	err = c.service.SetDoneStatus(inviteId, wasDone.Done)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	responses.Accepted(w, "Updated")
}

func (c InviteController) DeleteInvite(w http.ResponseWriter, r *http.Request, inviteId int) {
	err := c.service.RemoveInvite(inviteId)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	responses.NoContent(w, "Invite successfully deleted")
}
