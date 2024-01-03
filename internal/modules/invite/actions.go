package invite

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
	"github.com/andresmeireles/speaker/internal/logger"
	"github.com/andresmeireles/speaker/internal/modules/config"
	"github.com/andresmeireles/speaker/internal/modules/person"
)

func ParseInviteWithTemplate(
	inviteRepository repository.Repository[entity.Invite],
	configRepository config.ConfigRepository,
	inviteSender InviteSender,
) (string, error) {
	invite, err := inviteRepository.GetById(inviteSender.InvoiceId)
	if err != nil {
		return "", err
	}

	config, err := configRepository.GetByName("template")
	if err != nil {
		return "", err
	}

	template := config.Value
	inviteDate := invite.Date.Format("02/01/2006")
	replaceName := strings.Replace(template, "{{name}}", invite.Person.Name, -1)
	replaceDate := strings.Replace(replaceName, "{{date}}", inviteDate, -1)
	replaceTheme := strings.Replace(replaceDate, "{{theme}}", invite.Theme, -1)
	inviteText := strings.Replace(replaceTheme, "{{time}}", strconv.Itoa(invite.Time), -1)

	return inviteText, nil
}

func CreateInvite(
	inviteRepository repository.Repository[entity.Invite],
	personRepository repository.Repository[entity.Person],
	inviteData InvitePost,
) (entity.Invite, error) {
	personEntity, err := personRepository.GetById(inviteData.PersonId)
	if err != nil {
		logger.Error(err)

		return entity.Invite{}, fmt.Errorf("person with id %d not found", inviteData.PersonId)
	}

	iv := entity.Invite{
		PersonId: personEntity.GetId(),
		Theme:    inviteData.Theme,
		Time:     inviteData.Time,
		Date:     time.Now(),
	}
	err = inviteRepository.Add(iv)

	if err != nil {
		logger.Error(err)

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

func UpdateInvite(
	inviteRepository InviteRepository,
	personRepository person.PersonRepository,
	updateInviteData InvitePost,
	inviteId int,
) error {
	invite, err := inviteRepository.GetById(inviteId)
	if err != nil {
		logger.Error(err)

		return err
	}

	if invite.PersonId != updateInviteData.PersonId {
		person, err := personRepository.GetById(updateInviteData.PersonId)
		if err != nil {
			logger.Error(err)

			return err
		}

		invite.PersonId = person.GetId()
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
