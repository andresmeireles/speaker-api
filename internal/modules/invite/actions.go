package invite

import (
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
	"github.com/andresmeireles/speaker/internal/logger"
	"github.com/andresmeireles/speaker/internal/modules/config"
	"github.com/andresmeireles/speaker/internal/modules/person"
)

type Actions struct {
	inviteRepository InviteRepository
	personRepository person.PersonRepository
	configRepository config.ConfigRepository
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

func ParseRememberMessage(
	inviteRepository repository.Repository[entity.Invite],
	configRepository config.ConfigRepository,
	inviteId int,
) (string, error) {
	invite, err := inviteRepository.GetById(inviteId)
	if err != nil {
		slog.Error("error when get invite", err)

		return "", err
	}

	config, err := configRepository.GetByName("remember")
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
		slog.Error("Error", err)

		return entity.Invite{}, fmt.Errorf("person with id %d not found", inviteData.PersonId)
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
		logger.Error("error on delete invite, when get invite by id", id, err)

		return err
	}

	return repository.Delete(*invite)
}

func (a Actions) UpdateInvite(
	inviteRepository InviteRepository,
	personRepository person.PersonRepository,
	updateInviteData UpdateInviteData,
	inviteId int,
) error {
	invite, err := inviteRepository.GetById(inviteId)
	if err != nil {
		slog.Error("error when get id", err)

		return err
	}

	invite.Theme = updateInviteData.Theme
	invite.Time = updateInviteData.Time
	invite.Accepted = updateInviteData.Accepted
	invite.Remembered = updateInviteData.Remembered
	invite.References = updateInviteData.References

	return inviteRepository.Update(*invite)
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
		slog.Error("error on accept invite, when get invite by id", inviteId, err)

		return err
	}

	acceptQuery := "UPDATE invites SET accepted = true WHERE id = $1;"
	_, err = a.inviteRepository.Query(acceptQuery, inviteId)

	if err != nil {
		slog.Error("error on accept invite, when get invite by id;", inviteId, err)
	}

	return err
}

func (a Actions) rememberInvite(inviteId int) error {
	_, err := a.inviteRepository.GetById(inviteId)
	if err != nil {
		slog.Error("error on accept invite, when get invite by id", inviteId, err)

		return err
	}

	acceptQuery := "UPDATE invites SET remembered=true WHERE id = $1;"
	_, err = a.inviteRepository.Query(acceptQuery, inviteId)

	if err != nil {
		slog.Error("error on accept invite, when get invite by id;", inviteId, err)
	}

	return err
}
