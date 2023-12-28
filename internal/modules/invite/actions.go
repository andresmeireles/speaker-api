package invite

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
)

func ParseInviteWithTemplate(
	inviteRepository repository.Repository[entity.Invite],
	configRepository repository.Repository[entity.Config],
	inviteSender InviteSender,
) (string, error) {
	invite, err := inviteRepository.GetById(inviteSender.InvoiceId)
	if err != nil {
		return "", err
	}
	config, err := configRepository.GetById(inviteSender.TemplateId)
	if err != nil {
		return "", err
	}
	template := config.Value

	inviteDate := time.Now().Format("02/01/2006")
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
		return entity.Invite{}, errors.New(fmt.Sprintf("person with id %d not found", inviteData.PersonId))
	}

	iv := entity.Invite{
		Person: *personEntity,
		Theme:  inviteData.Theme,
		Time:   inviteData.Time,
		Date:   time.Now(),
	}
	err = inviteRepository.Add(iv)
	if err != nil {
		return entity.Invite{}, err
	}

	return iv, nil
}

func UpdateInvite(
	inviteRepository repository.Repository[entity.Invite],
	personRepository repository.Repository[entity.Person],
	inviteData InvitePost,
	inviteId int,
) (entity.Invite, error) {
	err := validateInviteData(inviteData)
	if err != nil {
		return entity.Invite{}, err
	}

	currentInvite, err := inviteRepository.GetById(inviteId)
	if err != nil {
		return entity.Invite{}, fmt.Errorf("invite with id %d not found", inviteId)
	}

	newPerson := &currentInvite.Person
	if inviteData.PersonId != currentInvite.Person.GetId() {
		newPerson, err = personRepository.GetById(inviteData.PersonId)
		if err != nil {
			return entity.Invite{}, errors.New(
				fmt.Sprintf("person with id %d not found", inviteData.PersonId),
			)
		}
		currentInvite.Person = *newPerson
	}

	currentInvite.Theme = inviteData.Theme
	currentInvite.Time = inviteData.Time
	err = inviteRepository.Update(*currentInvite)
	if err != nil {
		return entity.Invite{}, errors.New("could not update invite")
	}

	return *currentInvite, nil
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
