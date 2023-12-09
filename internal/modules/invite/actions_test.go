package invite_test

import (
	"errors"
	"testing"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/modules/invite"
	"github.com/stretchr/testify/mock"
)

type ConfigRepoMock struct {
	mock.Mock
}

func (m *ConfigRepoMock) GetById(id int) (*entity.Config, error) {
	args := m.Called(id)
	config := args.Get(0).(entity.Config)
	return &config, args.Error(1)
}

func (m *ConfigRepoMock) Add(config entity.Config) error {
	return nil
}

func (m *ConfigRepoMock) Update(config entity.Config) error {
	return nil
}

func (m *ConfigRepoMock) Delete(config entity.Config) error {
	return nil
}

func (m *ConfigRepoMock) GetAll() []entity.Config {
	return nil
}

func (m *ConfigRepoMock) GetByName(name string) (*entity.Config, error) {
	return nil, nil
}

type InviteRepoMock struct {
	mock.Mock
}

func (m *InviteRepoMock) GetById(id int) (*entity.Invite, error) {
	args := m.Called(id)
	invite := args.Get(0).(entity.Invite)
	return &invite, args.Error(1)
}

func (m *InviteRepoMock) Add(invite entity.Invite) error {
	return nil
}

func (m *InviteRepoMock) Update(invite entity.Invite) error {
	return nil
}

func (m *InviteRepoMock) Delete(invite entity.Invite) error {
	return nil
}

func (m *InviteRepoMock) GetAll() []entity.Invite {
	return nil
}

type PersonRepoMock struct {
	mock.Mock
}

func (m *PersonRepoMock) GetById(id int) (*entity.Person, error) {
	args := m.Called(id)
	person := args.Get(0).(entity.Person)
	return &person, args.Error(1)
}

func (m *PersonRepoMock) Add(person entity.Person) error {
	return nil
}
func (m *PersonRepoMock) Update(person entity.Person) error {
	return nil
}
func (m *PersonRepoMock) GetAll() []entity.Person {
	return nil
}
func (m *PersonRepoMock) Delete(person entity.Person) error {
	return nil
}

func TestActions(t *testing.T) {
	t.Run("should parse template by default", func(t *testing.T) {
		// arrange
		en := entity.Invite{
			Person: entity.Person{
				Name: "Person 1",
			},
			Theme:      "Theme",
			Time:       5,
			Date:       1166572800,
			Accepted:   true,
			Remembered: true,
		}
		template := "{{name}} invited you with theme {{theme}} with {{time}} minutes on {{date}}"

		inviteRepo := new(InviteRepoMock)
		configRepo := new(ConfigRepoMock)
		senderData := invite.InviteSender{
			InvoiceId:  1,
			TemplateId: 1,
		}

		inviteRepo.On("GetById", 1).Return(en, nil)
		configRepo.On("GetById", 1).Return(entity.Config{
			Id:    1,
			Value: template,
		}, nil)

		// act
		result, err := invite.ParseInviteWithTemplate(inviteRepo, configRepo, senderData)

		// assert
		if err != nil {
			t.Fatalf("expected no error, got %s", err)
		}
		if result != "Person 1 invited you with theme Theme with 5 minutes on 20/12/2006" {
			t.Fatalf("expected Person 1 invited you with theme Theme with 5 minutes on 20/12/2006, got %s", result)
		}
		inviteRepo.AssertNumberOfCalls(t, "GetById", 1)
		inviteRepo.AssertExpectations(t)
		configRepo.AssertNumberOfCalls(t, "GetById", 1)
		configRepo.AssertExpectations(t)
	})

	t.Run("should create new invite", func(t *testing.T) {
		// arrange
		invitePost := invite.InvitePost{
			PersonId: 1,
			Theme:    "Theme",
			Time:     5,
			Date:     "10/11/2024",
		}

		inviteRepo := new(InviteRepoMock)
		personRepo := new(PersonRepoMock)

		personRepo.On("GetById", 1).Return(entity.Person{
			Id:   1,
			Name: "Person 1",
		}, nil)

		// act
		result, err := invite.CreateInvite(
			inviteRepo,
			personRepo,
			invitePost,
		)

		// assert
		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}
		if result.Person.GetId() != 1 {
			t.Fatalf("expected 1, got %d", result.Person.GetId())
		}
		personRepo.AssertExpectations(t)
	})

	t.Run("should break when data is invalid", func(t *testing.T) {
		// arrange
		invitePost := invite.InvitePost{
			PersonId: 1,
			Theme:    "Theme",
			Time:     0,
			Date:     "bola",
		}

		inviteRepo := new(InviteRepoMock)

		personRepo := new(PersonRepoMock)

		personRepo.On("GetById", 1).Return(entity.Person{
			Id:   1,
			Name: "Person 1",
		}, nil)

		// act
		_, err := invite.CreateInvite(
			inviteRepo,
			personRepo,
			invitePost,
		)

		// assert
		if err == nil {
			t.Fatalf("this test should return an error")
		}
		if err.Error() != "invalid date" {
			t.Fatalf("expected invalid date, got %s", err)
		}
		personRepo.AssertExpectations(t)
		inviteRepo.AssertExpectations(t)
	})

	t.Run("should break when person is invalid", func(t *testing.T) {
		// arrange
		invitePost := invite.InvitePost{
			PersonId: 1,
			Theme:    "Theme",
			Time:     0,
			Date:     "bola",
		}
		inviteRepo := new(InviteRepoMock)
		personRepo := new(PersonRepoMock)
		personRepo.On("GetById", 1).Return(entity.Person{}, errors.New(""))

		// act
		_, err := invite.CreateInvite(
			inviteRepo,
			personRepo,
			invitePost,
		)

		// assert
		if err == nil {
			t.Fatalf("this test should return an error")
		}
		if err.Error() != "person with id 1 not found" {
			t.Fatalf("expected person with id 1 not found, got %s", err)
		}
		personRepo.AssertExpectations(t)
		inviteRepo.AssertExpectations(t)
	})

	t.Run("should break on 0 as time", func(t *testing.T) {
		// arrange
		invitePost := invite.InvitePost{
			PersonId: 1,
			Theme:    "Theme",
			Time:     0,
			Date:     "bola",
		}
		inviteRepo := new(InviteRepoMock)
		personRepo := new(PersonRepoMock)

		// act
		_, err := invite.CreateInvite(
			inviteRepo,
			personRepo,
			invitePost,
		)

		// assert
		if err == nil {
			t.Fatalf("this test should return an error")
		}
		if err.Error() != "invalid time, must be greater than 0" {
			t.Fatalf("expected invalid time, must be greater than 0, got %s", err)
		}
		personRepo.AssertNotCalled(t, "GetById")
		inviteRepo.AssertNotCalled(t, "Create")
	})

	t.Run("should break on empty date", func(t *testing.T) {
		// arrange
		invitePost := invite.InvitePost{
			PersonId: 1,
			Theme:    "Theme",
			Time:     200,
			Date:     "",
		}
		inviteRepo := new(InviteRepoMock)
		personRepo := new(PersonRepoMock)

		// act
		_, err := invite.CreateInvite(
			inviteRepo,
			personRepo,
			invitePost,
		)

		// assert
		if err.Error() != "invalid date, must be not empty" {
			t.Fatalf("expected invalid date, must be not empty, got %s", err)
		}
		personRepo.AssertNotCalled(t, "GetById")
		inviteRepo.AssertNotCalled(t, "Create")
	})

	t.Run("should break on empty theme", func(t *testing.T) {
		// arrange
		invitePost := invite.InvitePost{
			PersonId: 1,
			Theme:    "",
			Time:     200,
			Date:     "bola",
		}
		inviteRepo := new(InviteRepoMock)
		personRepo := new(PersonRepoMock)

		// act
		_, err := invite.CreateInvite(
			inviteRepo,
			personRepo,
			invitePost,
		)

		// assert
		if err.Error() != "invalid theme, must be not empty" {
			t.Fatalf("expected invalid theme, must be not empty, got %s", err)
		}
		personRepo.AssertNotCalled(t, "GetById")
		inviteRepo.AssertNotCalled(t, "Create")
	})
}
