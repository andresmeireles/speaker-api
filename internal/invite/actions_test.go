package invite_test

import (
	"testing"
	"time"

	"github.com/andresmeireles/speaker/internal/config"
	"github.com/andresmeireles/speaker/internal/invite"
	"github.com/andresmeireles/speaker/internal/person"
	"github.com/andresmeireles/speaker/testdata"
)

func TestParseTemplate(t *testing.T) {
	inviteRepo := testdata.GetService[invite.InviteRepository]()
	actions := testdata.GetService[invite.Actions]()
	personRepo := testdata.GetService[person.PersonRepository]()
	configRepo := testdata.GetService[config.ConfigRepository]()

	setupInvite := func() {
		err := personRepo.Add(person.Person{Name: "Person 1"})

		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}

		en := invite.Invite{
			PersonId:   1,
			Theme:      "Theme",
			Time:       5,
			Date:       time.Date(2006, 12, 20, 0, 0, 0, 0, time.UTC),
			References: "bola",
		}

		err = inviteRepo.Add(en)

		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}
	}

	t.Run("should parse template by default", func(t *testing.T) {
		// arrange
		setupInvite()
		template := "{{name}} invited you with theme {{theme}} with {{time}} minutes on {{date}}"
		err := configRepo.Add(config.Config{
			Name:  "template",
			Value: template,
		})

		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}

		invites, err := inviteRepo.GetAll()
		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}

		// act
		result, err := actions.ParseInviteWithTemplate(invites[0].Id)

		// assert
		if err != nil {
			t.Fatalf("expected no error, got %s", err)
		}

		if result != "Person 1 invited you with theme Theme with 5 minutes on 20/12/2006" {
			t.Fatalf("expected Person 1 invited you with theme Theme with 5 minutes on 20/12/2006, got %s", result)
		}
	})
}

func TestCreateInvite(t *testing.T) {
	personRepo := testdata.GetService[person.PersonRepository]()
	actions := testdata.GetService[invite.Actions]()

	t.Run("should create new invite", func(t *testing.T) {
		// arrange
		person := person.Person{
			Name: "Person 1000",
		}
		err := personRepo.Add(person)

		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}

		p, err := personRepo.GetByName(person.Name)
		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}

		invitePost := invite.InvitePost{
			PersonId: p.Id,
			Theme:    "Theme",
			Time:     5,
			Date:     time.Date(2006, 12, 20, 0, 0, 0, 0, time.UTC).Format("2006-01-02T15:04:05.000Z"),
		}

		// act
		result, err := actions.CreateInvite(invitePost)

		// assert
		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}

		if result.Theme != "Theme" {
			t.Fatalf("expected Theme, got %d", result.PersonId)
		}
	})
}

// 	t.Run("should break when data is invalid", func(t *testing.T) {
// 		// arrange
// 		invitePost := invite.InvitePost{
// 			PersonId: 1,
// 			Theme:    "Theme",
// 			Time:     0,
// 			Date:     "bola",
// 		}

// 		inviteRepo := new(InviteRepoMock)

// 		personRepo := new(PersonRepoMock)

// 		personRepo.On("GetById", 1).Return(entity.Person{
// 			Id:   1,
// 			Name: "Person 1",
// 		}, nil)

// 		// act
// 		_, err := invite.CreateInvite(
// 			inviteRepo,
// 			personRepo,
// 			invitePost,
// 		)

// 		// assert
// 		if err == nil {
// 			t.Fatalf("this test should return an error")
// 		}
// 		if err.Error() != "invalid date" {
// 			t.Fatalf("expected invalid date, got %s", err)
// 		}
// 		personRepo.AssertExpectations(t)
// 		inviteRepo.AssertExpectations(t)
// 	})

// 	t.Run("should break when person is invalid", func(t *testing.T) {
// 		// arrange
// 		invitePost := invite.InvitePost{
// 			PersonId: 1,
// 			Theme:    "Theme",
// 			Time:     0,
// 			Date:     "bola",
// 		}
// 		inviteRepo := new(InviteRepoMock)
// 		personRepo := new(PersonRepoMock)
// 		personRepo.On("GetById", 1).Return(entity.Person{}, errors.New(""))

// 		// act
// 		_, err := invite.CreateInvite(
// 			inviteRepo,
// 			personRepo,
// 			invitePost,
// 		)

// 		// assert
// 		if err == nil {
// 			t.Fatalf("this test should return an error")
// 		}
// 		if err.Error() != "person with id 1 not found" {
// 			t.Fatalf("expected person with id 1 not found, got %s", err)
// 		}
// 		personRepo.AssertExpectations(t)
// 		inviteRepo.AssertExpectations(t)
// 	})

// 	t.Run("should break on 0 as time", func(t *testing.T) {
// 		// arrange
// 		invitePost := invite.InvitePost{
// 			PersonId: 1,
// 			Theme:    "Theme",
// 			Time:     0,
// 			Date:     "bola",
// 		}
// 		inviteRepo := new(InviteRepoMock)
// 		personRepo := new(PersonRepoMock)

// 		// act
// 		_, err := invite.CreateInvite(
// 			inviteRepo,
// 			personRepo,
// 			invitePost,
// 		)

// 		// assert
// 		if err == nil {
// 			t.Fatalf("this test should return an error")
// 		}
// 		if err.Error() != "invalid time, must be greater than 0" {
// 			t.Fatalf("expected invalid time, must be greater than 0, got %s", err)
// 		}
// 		personRepo.AssertNotCalled(t, "GetById")
// 		inviteRepo.AssertNotCalled(t, "Create")
// 	})

// 	t.Run("should break on empty date", func(t *testing.T) {
// 		// arrange
// 		invitePost := invite.InvitePost{
// 			PersonId: 1,
// 			Theme:    "Theme",
// 			Time:     200,
// 			Date:     "",
// 		}
// 		inviteRepo := new(InviteRepoMock)
// 		personRepo := new(PersonRepoMock)

// 		// act
// 		_, err := invite.CreateInvite(
// 			inviteRepo,
// 			personRepo,
// 			invitePost,
// 		)

// 		// assert
// 		if err.Error() != "invalid date, must be not empty" {
// 			t.Fatalf("expected invalid date, must be not empty, got %s", err)
// 		}
// 		personRepo.AssertNotCalled(t, "GetById")
// 		inviteRepo.AssertNotCalled(t, "Create")
// 	})

// 	t.Run("should break on empty theme", func(t *testing.T) {
// 		// arrange
// 		invitePost := invite.InvitePost{
// 			PersonId: 1,
// 			Theme:    "",
// 			Time:     200,
// 			Date:     "bola",
// 		}
// 		inviteRepo := new(InviteRepoMock)
// 		personRepo := new(PersonRepoMock)

// 		// act
// 		_, err := invite.CreateInvite(
// 			inviteRepo,
// 			personRepo,
// 			invitePost,
// 		)

//		// assert
//		if err.Error() != "invalid theme, must be not empty" {
//			t.Fatalf("expected invalid theme, must be not empty, got %s", err)
//		}
//		personRepo.AssertNotCalled(t, "GetById")
//		inviteRepo.AssertNotCalled(t, "Create")
//	})
// }
