package invite_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/modules/invite"
	"github.com/andresmeireles/speaker/internal/modules/person"
	"github.com/andresmeireles/speaker/testdata"
)

func TestMain(m *testing.M) {
	testdata.SetupDatabase(m)
}

func TestController(t *testing.T) {
	t.Run("should execute creation", func(t *testing.T) {
		// arrange
		personRepo := person.PersonRepository{}
		err := personRepo.Add(entity.Person{Name: "Andre"})

		if err != nil {
			t.Fatalf("expected nil, got %s", err)
		}

		controller := testdata.GetService[invite.InviteController]()

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
