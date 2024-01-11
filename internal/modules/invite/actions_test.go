package invite_test

// import (
// 	"testing"
// 	"time"

// 	"github.com/andresmeireles/speaker/internal/db/entity"
// 	"github.com/andresmeireles/speaker/internal/modules/config"
// 	"github.com/andresmeireles/speaker/internal/modules/invite"
// 	"github.com/andresmeireles/speaker/internal/modules/person"
// )

// func TestParseTemplate(t *testing.T) {
// 	inviteRepo := invite.InviteRepository{}
// 	actions := invite.NewActions()

// 	setupInvoice := func() {
// 		personRepo := person.PersonRepository{}
// 		err := personRepo.Add(entity.Person{Name: "Person 1"})

// 		if err != nil {
// 			t.Fatalf("expected nil, got %s", err)
// 		}

// 		en := entity.Invite{
// 			PersonId:   1,
// 			Theme:      "Theme",
// 			Time:       5,
// 			Date:       time.Date(2006, 12, 20, 0, 0, 0, 0, time.UTC),
// 			Accepted:   true,
// 			Remembered: true,
// 			References: "bola",
// 		}

// 		err = inviteRepo.Add(en)

// 		if err != nil {
// 			t.Fatalf("expected nil, got %s", err)
// 		}
// 	}

// 	t.Run("should parse template by default", func(t *testing.T) {
// 		// arrange
// 		setupInvoice()
// 		template := "{{name}} invited you with theme {{theme}} with {{time}} minutes on {{date}}"
// 		configRepo := config.ConfigRepository{}
// 		err := configRepo.Add(entity.Config{
// 			Name:  "template",
// 			Value: template,
// 		})

// 		if err != nil {
// 			t.Fatalf("expected nil, got %s", err)
// 		}

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
// }

// func TestCreateInvite(t *testing.T) {
// 	personRepo := person.PersonRepository{}
// 	actions := invite.NewActions()

// 	t.Run("should create new invite", func(t *testing.T) {
// 		// arrange
// 		person := entity.Person{
// 			Name: "Person 1",
// 		}
// 		personRepo.Add(person)
// 		invitePost := invite.InvitePost{
// 			PersonId: 1,
// 			Theme:    "Theme",
// 			Time:     5,
// 			Date:     time.Date(2006, 12, 20, 0, 0, 0, 0, time.UTC).String(),
// 		}

// 		// act
// 		result, err := actions.CreateInvite(invitePost)

// 		// assert
// 		if err != nil {
// 			t.Fatalf("expected nil, got %s", err)
// 		}

// 		if result.PersonId != 1 {
// 			t.Fatalf("expected 1, got %d", result.PersonId)
// 		}
// 	})

// 	// 	t.Run("should break when data is invalid", func(t *testing.T) {
// 	// 		// arrange
// 	// 		invitePost := invite.InvitePost{
// 	// 			PersonId: 1,
// 	// 			Theme:    "Theme",
// 	// 			Time:     0,
// 	// 			Date:     "bola",
// 	// 		}

// 	// 		inviteRepo := new(InviteRepoMock)

// 	// 		personRepo := new(PersonRepoMock)

// 	// 		personRepo.On("GetById", 1).Return(entity.Person{
// 	// 			Id:   1,
// 	// 			Name: "Person 1",
// 	// 		}, nil)

// 	// 		// act
// 	// 		_, err := invite.CreateInvite(
// 	// 			inviteRepo,
// 	// 			personRepo,
// 	// 			invitePost,
// 	// 		)

// 	// 		// assert
// 	// 		if err == nil {
// 	// 			t.Fatalf("this test should return an error")
// 	// 		}
// 	// 		if err.Error() != "invalid date" {
// 	// 			t.Fatalf("expected invalid date, got %s", err)
// 	// 		}
// 	// 		personRepo.AssertExpectations(t)
// 	// 		inviteRepo.AssertExpectations(t)
// 	// 	})

// 	// 	t.Run("should break when person is invalid", func(t *testing.T) {
// 	// 		// arrange
// 	// 		invitePost := invite.InvitePost{
// 	// 			PersonId: 1,
// 	// 			Theme:    "Theme",
// 	// 			Time:     0,
// 	// 			Date:     "bola",
// 	// 		}
// 	// 		inviteRepo := new(InviteRepoMock)
// 	// 		personRepo := new(PersonRepoMock)
// 	// 		personRepo.On("GetById", 1).Return(entity.Person{}, errors.New(""))

// 	// 		// act
// 	// 		_, err := invite.CreateInvite(
// 	// 			inviteRepo,
// 	// 			personRepo,
// 	// 			invitePost,
// 	// 		)

// 	// 		// assert
// 	// 		if err == nil {
// 	// 			t.Fatalf("this test should return an error")
// 	// 		}
// 	// 		if err.Error() != "person with id 1 not found" {
// 	// 			t.Fatalf("expected person with id 1 not found, got %s", err)
// 	// 		}
// 	// 		personRepo.AssertExpectations(t)
// 	// 		inviteRepo.AssertExpectations(t)
// 	// 	})

// 	// 	t.Run("should break on 0 as time", func(t *testing.T) {
// 	// 		// arrange
// 	// 		invitePost := invite.InvitePost{
// 	// 			PersonId: 1,
// 	// 			Theme:    "Theme",
// 	// 			Time:     0,
// 	// 			Date:     "bola",
// 	// 		}
// 	// 		inviteRepo := new(InviteRepoMock)
// 	// 		personRepo := new(PersonRepoMock)

// 	// 		// act
// 	// 		_, err := invite.CreateInvite(
// 	// 			inviteRepo,
// 	// 			personRepo,
// 	// 			invitePost,
// 	// 		)

// 	// 		// assert
// 	// 		if err == nil {
// 	// 			t.Fatalf("this test should return an error")
// 	// 		}
// 	// 		if err.Error() != "invalid time, must be greater than 0" {
// 	// 			t.Fatalf("expected invalid time, must be greater than 0, got %s", err)
// 	// 		}
// 	// 		personRepo.AssertNotCalled(t, "GetById")
// 	// 		inviteRepo.AssertNotCalled(t, "Create")
// 	// 	})

// 	// 	t.Run("should break on empty date", func(t *testing.T) {
// 	// 		// arrange
// 	// 		invitePost := invite.InvitePost{
// 	// 			PersonId: 1,
// 	// 			Theme:    "Theme",
// 	// 			Time:     200,
// 	// 			Date:     "",
// 	// 		}
// 	// 		inviteRepo := new(InviteRepoMock)
// 	// 		personRepo := new(PersonRepoMock)

// 	// 		// act
// 	// 		_, err := invite.CreateInvite(
// 	// 			inviteRepo,
// 	// 			personRepo,
// 	// 			invitePost,
// 	// 		)

// 	// 		// assert
// 	// 		if err.Error() != "invalid date, must be not empty" {
// 	// 			t.Fatalf("expected invalid date, must be not empty, got %s", err)
// 	// 		}
// 	// 		personRepo.AssertNotCalled(t, "GetById")
// 	// 		inviteRepo.AssertNotCalled(t, "Create")
// 	// 	})

// 	// 	t.Run("should break on empty theme", func(t *testing.T) {
// 	// 		// arrange
// 	// 		invitePost := invite.InvitePost{
// 	// 			PersonId: 1,
// 	// 			Theme:    "",
// 	// 			Time:     200,
// 	// 			Date:     "bola",
// 	// 		}
// 	// 		inviteRepo := new(InviteRepoMock)
// 	// 		personRepo := new(PersonRepoMock)

// 	// 		// act
// 	// 		_, err := invite.CreateInvite(
// 	// 			inviteRepo,
// 	// 			personRepo,
// 	// 			invitePost,
// 	// 		)

// 	//		// assert
// 	//		if err.Error() != "invalid theme, must be not empty" {
// 	//			t.Fatalf("expected invalid theme, must be not empty, got %s", err)
// 	//		}
// 	//		personRepo.AssertNotCalled(t, "GetById")
// 	//		inviteRepo.AssertNotCalled(t, "Create")
// 	//	})
// }
