package invite_test

// import (
// 	"errors"
// 	"testing"
// 	"time"

// 	"github.com/andresmeireles/speaker/internal/config"
// 	"github.com/andresmeireles/speaker/internal/invite"
// 	"github.com/andresmeireles/speaker/internal/person"
// 	"github.com/andresmeireles/speaker/testdata/mocks"
// )

// //nolint:funlen,gocognit
// func TestParseTemplate(t *testing.T) {
// 	inviteRepoMock := mocks.InviteRepository{}
// 	personRepoMock := mocks.PersonRepository{}
// 	configRepoMock := mocks.ConfigRepo{}

// 	actions := invite.NewAction(
// 		&inviteRepoMock,
// 		&personRepoMock,
// 		&configRepoMock,
// 	)

// 	t.Run("should parse template by default", func(t *testing.T) {
// 		// arrange
// 		inviteRepoMock.EXPECT().GetById(1).Return(&invite.Invite{
// 			Person: person.Person{Name: "Person 1"},
// 			Theme:  "Theme",
// 			Time:   5,
// 			Date:   time.Date(2006, 12, 20, 0, 0, 0, 0, time.UTC),
// 		}, nil).Once()
// 		configRepoMock.EXPECT().GetByName("template").Return(&config.Config{
// 			Value: "{{name}} invited you with theme {{theme}} with {{time}} minutes on {{date}}",
// 		}, nil).Once()

// 		// act
// 		result, err := actions.ParseInviteWithTemplate(1)

// 		// assert
// 		if err != nil {
// 			t.Fatalf("expected no error, got %s", err)
// 		}

// 		if result != "Person 1 invited you with theme Theme with 5 minutes on 20/12/2006" {
// 			t.Fatalf("expected Person 1 invited you with theme Theme with 5 minutes on 20/12/2006, got %s", result)
// 		}
// 	})

// 	t.Run("set done status as done", func(t *testing.T) {
// 		// arrange
// 		inviteRepoMock.EXPECT().GetById(1).Return(&invite.Invite{}, nil).Once()
// 		inviteRepoMock.EXPECT().UpdateStatus(invite.Invite{}, invite.STATUS_DONE).Return(nil).Once()

// 		// act
// 		err := actions.SetDoneStatus(1, true)

// 		// assert
// 		if err != nil {
// 			t.Fatalf("expected nil, got %s", err)
// 		}
// 	})

// 	t.Run("set done status as not done", func(t *testing.T) {
// 		// arrange
// 		inviteRepoMock.EXPECT().GetById(1).Return(&invite.Invite{}, nil).Once()
// 		inviteRepoMock.EXPECT().UpdateStatus(invite.Invite{}, invite.STATUS_NOT_DONE).Return(nil).Once()

// 		// act
// 		err := actions.SetDoneStatus(1, false)

// 		// assert
// 		if err != nil {
// 			t.Fatalf("expected nil, got %s", err)
// 		}
// 	})

// 	t.Run("should create new invite", func(t *testing.T) {
// 		// arrange
// 		invitePost := invite.InvitePost{
// 			PersonId: 1,
// 			Theme:    "Theme",
// 			Time:     5,
// 			Date:     "2006-12-20T00:00:00.000Z",
// 		}

// 		personRepoMock.EXPECT().GetById(1).Return(&person.Person{
// 			Name: "Person 1",
// 		}, nil).Once()

// 		inviteRepoMock.EXPECT().Add(invite.Invite{
// 			Person: person.Person{Name: ""},
// 			Theme:  "Theme",
// 			Time:   5,
// 			Date:   time.Date(2006, 12, 20, 0, 0, 0, 0, time.UTC),
// 		}).Return(nil).Once()

// 		// act
// 		result, err := actions.CreateInvite(invitePost)

// 		// assert
// 		if err != nil {
// 			t.Fatalf("expected nil, got %s", err)
// 		}

// 		if result.Theme != "Theme" {
// 			t.Fatalf("expected Theme, got %d", result.PersonId)
// 		}
// 	})

// 	t.Run("should break when data is invalid", func(t *testing.T) {
// 		// arrange
// 		invitePost := invite.InvitePost{
// 			PersonId: 1,
// 			Theme:    "Theme",
// 			Time:     3,
// 			Date:     "bola",
// 		}

// 		personRepoMock.EXPECT().GetById(1).Return(&person.Person{}, nil).Once()
// 		inviteRepoMock.EXPECT().Add(invite.Invite{}).Return(errors.New("invalid date")).Once()

// 		// act
// 		_, err := actions.CreateInvite(invitePost)

// 		// assert
// 		if err == nil {
// 			t.Fatalf("this test should return an error")
// 		}
// 		if err.Error() != "parsing time \"bola\" as \"2006-01-02T15:04:05.000Z\": cannot parse \"bola\" as \"2006\"" {
// 			t.Fatalf("expected invalid date, got %s", err)
// 		}
// 	})

// 	t.Run("should break when person is invalid", func(t *testing.T) {
// 		// arrange
// 		invitePost := invite.InvitePost{
// 			PersonId: 1,
// 			Theme:    "Theme",
// 			Time:     4,
// 			Date:     "bola",
// 		}

// 		personRepoMock.EXPECT().GetById(1).Return(nil, errors.New("person with id 1 not found")).Once()
// 		inviteRepoMock.EXPECT().Add(invite.Invite{}).Return(nil).Once()

// 		// act
// 		_, err := actions.CreateInvite(invitePost)

// 		// assert
// 		if err == nil {
// 			t.Fatalf("this test should return an error")
// 		}
// 		if err.Error() != "person with id 1 not found" {
// 			t.Fatalf("expected person with id 1 not found, got %s", err)
// 		}
// 	})

// 	t.Run("should break on 0 as time", func(t *testing.T) {
// 		// arrange
// 		invitePost := invite.InvitePost{
// 			PersonId: 1,
// 			Theme:    "Theme",
// 			Time:     0,
// 			Date:     "bola",
// 		}

// 		personRepoMock.EXPECT().GetById(1).Return(&person.Person{}, nil).Once()

// 		// act
// 		_, err := actions.CreateInvite(invitePost)

// 		// assert
// 		if err == nil {
// 			t.Fatalf("this test should return an error")
// 		}
// 		if err.Error() != "time must be greater than 0" {
// 			t.Fatalf("expected time must be greater than 0, got %s", err)
// 		}
// 	})

// 	t.Run("should break on negative time", func(t *testing.T) {
// 		// arrange
// 		invitePost := invite.InvitePost{
// 			PersonId: 1,
// 			Theme:    "Theme",
// 			Time:     -2,
// 			Date:     "bola",
// 		}

// 		personRepoMock.EXPECT().GetById(1).Return(&person.Person{}, nil).Once()

// 		// act
// 		_, err := actions.CreateInvite(invitePost)

// 		// assert
// 		if err == nil {
// 			t.Fatalf("this test should return an error")
// 		}
// 		if err.Error() != "time must be greater than 0" {
// 			t.Fatalf("expected time must be greater than 0, got %s", err)
// 		}
// 	})

// 	t.Run("should break on empty date", func(t *testing.T) {
// 		// arrange
// 		invitePost := invite.InvitePost{
// 			PersonId: 1,
// 			Theme:    "Theme",
// 			Time:     200,
// 			Date:     "",
// 		}

// 		personRepoMock.EXPECT().GetById(1).Return(&person.Person{}, nil).Once()

// 		// act
// 		_, err := actions.CreateInvite(invitePost)

// 		// assert
// 		if err.Error() != "parsing time \"\" as \"2006-01-02T15:04:05.000Z\": cannot parse \"\" as \"2006\"" {
// 			t.Fatalf("expected invalid date, must be not empty, got %s", err)
// 		}
// 	})

// 	t.Run("should break on empty theme", func(t *testing.T) {
// 		// arrange
// 		invitePost := invite.InvitePost{
// 			PersonId: 1,
// 			Theme:    "",
// 			Time:     200,
// 			Date:     "bola",
// 		}

// 		// act
// 		_, err := actions.CreateInvite(invitePost)

// 		// assert
// 		if err.Error() != "theme must be not empty" {
// 			t.Fatalf("expected theme must be not empty, got %s", err)
// 		}
// 	})
// }
