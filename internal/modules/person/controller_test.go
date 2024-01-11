package person_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/modules/person"
	"github.com/andresmeireles/speaker/testdata"
)

func TestController(t *testing.T) {
	controller := testdata.GetService[person.PersonController]()
	repo := testdata.GetService[person.PersonRepository]()

	t.Run("should write a new person", func(t *testing.T) {
		cleanDb()

		// arrange
		reader := strings.NewReader(`{"name":"Andre"}`)
		request, err := http.NewRequest(http.MethodPost, "/speakers", reader)

		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		// act
		handler := http.HandlerFunc(controller.WritePerson)
		handler.ServeHTTP(recorder, request)

		// assert
		if status := recorder.Code; status != http.StatusCreated {
			t.Fatalf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		if recorder.Body.String() != "Person created" {
			t.Errorf("handler returned unexpected body: got %v want %v", recorder.Body.String(), "Person created")
		}
	})

	t.Run("should rewrite config", func(t *testing.T) {
		// arrange
		cleanDb()
		repo.Add(entity.Person{Name: "Andre"})
		p, e := repo.GetAll()

		if e != nil {
			t.Fatalf("expected nil, got %s", e)
		}

		l := p[len(p)-1]
		id := strconv.Itoa(l.Id)
		reader := strings.NewReader(`{"speaker": "` + id + `"}`)
		recorder := httptest.NewRecorder()
		request, err := http.NewRequest(http.MethodDelete, "/speakers", reader)
		handler := http.HandlerFunc(controller.DeletePerson)

		if err != nil {
			t.Fatal(err)
		}

		// act
		handler.ServeHTTP(recorder, request)

		// assert
		if status := recorder.Code; status != 202 {
			t.Fatalf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		if recorder.Body.String() != "removed person" {
			t.Errorf("handler returned unexpected body: got %v want %v", recorder.Body.String(), "Person created")
		}
	})
}
