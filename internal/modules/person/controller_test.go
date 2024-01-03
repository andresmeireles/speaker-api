package person_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/modules/person"
)

func TestPersonController(t *testing.T) {
	t.Run("should write a new person", func(t *testing.T) {
		// arrange
		reader := strings.NewReader(`{"name":"Andre"}`)
		request, err := http.NewRequest("POST", "/speakers", reader)

		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		// act
		handler := http.HandlerFunc(person.WritePerson)
		handler.ServeHTTP(recorder, request)

		// assert
		if status := recorder.Code; status != http.StatusCreated {
			t.Fatalf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		if recorder.Body.String() != "Person created" {
			t.Errorf("handler returned unexpected body: got %v want %v", recorder.Body.String(), "Person created")
		}
	})

	t.Run("should remove person", func(t *testing.T) {
		// arrange
		person.PersonRepository{}.Add(entity.Person{Name: "Andre"})
		reader := strings.NewReader(`{"speaker": "1"}`)
		recorder := httptest.NewRecorder()
		request, err := http.NewRequest("DELETE", "/speakers", reader)
		handler := http.HandlerFunc(person.DeletePerson)

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
