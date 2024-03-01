package invite_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/andresmeireles/speaker/internal/invite"
	"github.com/andresmeireles/speaker/testdata/mocks"
)

func TestController(t *testing.T) {
	inviteRepoMock := mocks.InviteRepositoryinvite{}
	personRepoMock := mocks.PersonRepositoryperson{}
	configRepoMock := mocks.ConfigRepositoryconfig{}
	inviteServiceMock := mocks.InviteServiceinvite{}

	controller := invite.NewController(
		&inviteRepoMock,
		&personRepoMock,
		&configRepoMock,
		&inviteServiceMock,
	)

	t.Run("should execute creation", func(t *testing.T) {
		// arrange
		inviteServiceMock.EXPECT().
			CreateInvite(invite.InvitePost{PersonId: 1, Date: "2006-01-02T15:04:05.000Z", Theme: "Theme", Time: 1}).
			Return(invite.Invite{}, nil).
			Once()

		reader := strings.NewReader(`{"person_id":1,"date":"2006-01-02T15:04:05.000Z","theme":"Theme","time":1}`)
		recorder := httptest.NewRecorder()
		request, err := http.NewRequest(http.MethodPost, "/invites", reader)
		handler := http.HandlerFunc(controller.Create)

		if err != nil {
			t.Fatal(err)
		}

		// act
		handler.ServeHTTP(recorder, request)

		// assert
		if status := recorder.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if recorder.Body.String() != "Invite successfully created" {
			t.Errorf("handler returned unexpected body: got %v want %v",
				recorder.Body.String(), "Invite successfully created")
		}
	})
}
