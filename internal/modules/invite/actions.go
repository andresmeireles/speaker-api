package invite

import (
	"errors"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/modules/config"
	"github.com/andresmeireles/speaker/internal/modules/person"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
)

type Actions struct {
	inviteRepository InviteRepository
	personRepository person.PersonRepository
	configRepository config.ConfigRepository
}

func (a Actions) New(s servicelocator.ServiceLocator) any {
	inviteRepository := servicelocator.Get[InviteRepository](s)
	personRepository := servicelocator.Get[person.PersonRepository](s)
	configRepository := servicelocator.Get[config.ConfigRepository](s)

	return Actions{
		inviteRepository: inviteRepository,
		personRepository: personRepository,
		configRepository: configRepository,
	}
}

func (a Actions) ParseInviteWithTemplate(inviteId int) (string, error) {
	invite, err := a.inviteRepository.GetById(inviteId)
	if err != nil {
		slog.Error("error when get invite", err)

		return "", err
	}

	config, err := a.configRepository.GetByName("template")
	if err != nil {
		slog.Error("error when get config", err)

		return "", err
	}

	inviteText := parseMessage(config.Value, *invite)

	return inviteText, nil
}

func (a Actions) ParseRememberMessage(inviteId int) (string, error) {
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

func parseMessage(message string, invite entity.Invite) string {
	inviteDate := invite.Date.Format("02/01/2006")
	parsedMessage := strings.Replace(message, "{{name}}", invite.Person.Name, -1)
	parsedMessage = strings.Replace(parsedMessage, "{{date}}", inviteDate, -1)
	parsedMessage = strings.Replace(parsedMessage, "{{theme}}", invite.Theme, -1)
	parsedMessage = strings.Replace(parsedMessage, "{{time}}", strconv.Itoa(invite.Time), -1)
	parsedMessage = strings.Replace(parsedMessage, "{{references}}", invite.References, -1)

	return parsedMessage
}

func (a Actions) CreateInvite(
	inviteData InvitePost,
) (entity.Invite, error) {
	personEntity, err := a.personRepository.GetById(inviteData.PersonId)
	if err != nil {
		slog.Error("Error on get person", err)

		return entity.Invite{}, err
	}

	layout := "2006-01-02T15:04:05.000Z"
	date, err := time.Parse(layout, inviteData.Date)

	if err != nil {
		slog.Error("Error on parse", err)

		return entity.Invite{}, err
	}

	iv := entity.Invite{
		PersonId:   personEntity.GetId(),
		Theme:      inviteData.Theme,
		Time:       inviteData.Time,
		Date:       date,
		References: inviteData.References,
	}
	err = a.inviteRepository.Add(iv)

	if err != nil {
		slog.Error("Error", err)

		return entity.Invite{}, err
	}

	return iv, nil
}

func RemoveInvite(id int, repository InviteRepository) error {
	invite, err := repository.GetById(id)
	if err != nil {
		slog.Error("error on delete invite, when get invite by id", "invite id", id, err)

		return err
	}

	return repository.Delete(*invite)
}

func (a Actions) UpdateInvite(
	updateInviteData UpdateInviteData,
	inviteId int,
) error {
	invite, err := a.inviteRepository.GetById(inviteId)
	if err != nil {
		slog.Error("error when get id", err)

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

func (a Actions) acceptInvite(inviteId int) error {
	_, err := a.inviteRepository.GetById(inviteId)
	if err != nil {
		slog.Error("error on accept invite, when get invite by id", "invite id", inviteId, err)

		return err
	}

	acceptQuery := "UPDATE invites SET accepted = true WHERE id = $1;"
	_, err = a.inviteRepository.Query(acceptQuery, inviteId)

	if err != nil {
		slog.Error("error on accept invite, when get invite by id;", "invite id", inviteId, err)
	}

	return err
}

func (a Actions) rememberInvite(inviteId int) error {
	_, err := a.inviteRepository.GetById(inviteId)
	if err != nil {
		slog.Error("error on accept invite, when get invite by id", "invite id", inviteId, err)

		return err
	}

	acceptQuery := "UPDATE invites SET remembered=true WHERE id = $1;"
	_, err = a.inviteRepository.Query(acceptQuery, inviteId)

	if err != nil {
		slog.Error("error on accept invite, when get invite by id;", "invite id", inviteId, err)
	}

	return err
}
