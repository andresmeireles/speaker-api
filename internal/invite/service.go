package invite

import (
	"errors"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/andresmeireles/speaker/internal/config"
	"github.com/andresmeireles/speaker/internal/person"
)

type InviteService interface {
	ParseInviteWithTemplate(inviteId int) (string, error)
	ParseRememberMessage(inviteId int) (string, error)
	CreateInvite(inviteData InvitePost) (Invite, error)
	RemoveInvite(id int) error
	AcceptInvite(inviteId int) error
	RememberInvite(inviteId int) error
	SetDoneStatus(inviteId int, done bool) error
	UpdateInvite(updateInviteData UpdateInviteData, inviteId int) error
	Reject(inviteId int) error
}

type Service struct {
	inviteRepository InviteRepository
	personRepository person.PersonRepository
	configRepository config.ConfigRepository
}

func NewAction(
	inviteRepository InviteRepository,
	personRepository person.PersonRepository,
	configRepository config.ConfigRepository,
) Service {
	return Service{
		inviteRepository: inviteRepository,
		personRepository: personRepository,
		configRepository: configRepository,
	}
}

func (a Service) ParseInviteWithTemplate(inviteId int) (string, error) {
	invite, err := a.inviteRepository.GetById(inviteId)
	if err != nil {
		return "", err
	}

	config, err := a.configRepository.GetByName("template")
	if err != nil {
		return "", err
	}

	inviteText := parseMessage(config.Value, *invite)

	return inviteText, nil
}

// parseRememberMessage parses the remember message.
//
// Parameter(s): inviteId int
//
// Return type(s):
//
// string
// error
func (a Service) ParseRememberMessage(inviteId int) (string, error) {
	invite, err := a.inviteRepository.GetById(inviteId)
	if err != nil {
		slog.Error("error when get invite", err)

		return "", err
	}

	config, err := a.configRepository.GetByName("remember")
	if err != nil {
		slog.Error("error when get config", err)

		return "", err
	}

	parseMessage := parseMessage(config.Value, *invite)

	return parseMessage, nil
}

func parseMessage(message string, invite Invite) string {
	inviteDate := invite.Date.Format("02/01/2006")
	parsedMessage := strings.Replace(message, "{{name}}", invite.Person.Name, -1)
	parsedMessage = strings.Replace(parsedMessage, "{{date}}", inviteDate, -1)
	parsedMessage = strings.Replace(parsedMessage, "{{theme}}", invite.Theme, -1)
	parsedMessage = strings.Replace(parsedMessage, "{{time}}", strconv.Itoa(invite.Time), -1)
	parsedMessage = strings.Replace(parsedMessage, "{{references}}", invite.References, -1)

	return parsedMessage
}

// CreateInvite creates an invite using the provided invite data.
//
// Parameter(s):
//
//	inviteData InvitePost
//
// Return type(s):
//
//	Invite
//	error
func (a Service) CreateInvite(inviteData InvitePost) (Invite, error) {
	if inviteData.Theme == "" {
		return Invite{}, errors.New("theme must be not empty")
	}

	if inviteData.Time <= 0 {
		return Invite{}, errors.New("time must be greater than 0")
	}

	personEntity, err := a.personRepository.GetById(inviteData.PersonId)
	if err != nil {
		return Invite{}, err
	}

	layout := "2006-01-02T15:04:05.000Z"
	date, err := time.Parse(layout, inviteData.Date)

	if err != nil {
		return Invite{}, err
	}

	iv := Invite{
		PersonId:   personEntity.GetId(),
		Theme:      inviteData.Theme,
		Time:       inviteData.Time,
		Date:       date,
		References: inviteData.References,
	}
	err = a.inviteRepository.Add(iv)

	if err != nil {
		return Invite{}, err
	}

	return iv, nil
}

func (a Service) RemoveInvite(id int) error {
	invite, err := a.inviteRepository.GetById(id)
	if err != nil {
		return err
	}

	return a.inviteRepository.Delete(*invite)
}

func (a Service) UpdateInvite(
	updateInviteData UpdateInviteData,
	inviteId int,
) error {
	invite, err := a.inviteRepository.GetById(inviteId)
	if err != nil {
		return err
	}

	// string to time golang
	layout := "2006-01-02T15:04:05.000Z"
	date, err := time.Parse(layout, updateInviteData.Date)

	if err != nil {
		slog.Error("Error on parse", err)
	}

	invite.Theme = updateInviteData.Theme
	invite.Time = updateInviteData.Time
	invite.Date = date
	invite.References = updateInviteData.References

	return a.inviteRepository.Update(*invite)
}

func validateInviteData(inviteData InvitePost) error {
	if inviteData.Time == 0 {
		return errors.New("invalid time, must be greater than 0")
	}

	if len(strings.Trim(inviteData.Date, "")) == 0 {
		return errors.New("invalid date, must be not empty")
	}

	if len(strings.Trim(inviteData.Theme, "")) == 0 {
		return errors.New("invalid theme, must be not empty")
	}

	return nil
}

func (a Service) SetDoneStatus(inviteId int, done bool) error {
	d := STATUS_DONE
	if !done {
		d = STATUS_NOT_DONE
	}

	return a.updateStatus(inviteId, d)
}

func (a Service) AcceptInvite(inviteId int) error {
	return a.updateStatus(inviteId, STATUS_CONFIRMED)
}

func (a Service) RememberInvite(inviteId int) error {
	return a.updateStatus(inviteId, STATUS_REMEMBERED)
}

func (a Service) Reject(inviteId int) error {
	return a.updateStatus(inviteId, STATUS_REJECTED)
}

// updateStatus updates the status of an invite.
//
// Parameters:
//
//	inviteId int - the ID of the invite
//	status int - the new status to update
//
// Return type:
//
// error
func (a Service) updateStatus(inviteId int, status int) error {
	invite, err := a.inviteRepository.GetById(inviteId)
	if err != nil {
		return err
	}

	return a.inviteRepository.UpdateStatus(*invite, status)
}
