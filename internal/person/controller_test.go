package person_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/andresmeireles/speaker/internal/person"
	"github.com/andresmeireles/speaker/testdata"
	"github.com/stretchr/testify/mock"
)

func TestController(t *testing.T) {
	repositoryMock := testdata.PersonRepositoryMock{}
	actionsMock := testdata.PersonActionMock{}

	controller := person.NewController(&repositoryMock, &actionsMock)

	t.Run("should write a new person", func(t *testing.T) {
		// arrange
		actionsMock.On("Write", mock.Anything).Return(nil)
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
		reader := strings.NewReader(`{"speaker": "1"}`)
		recorder := httptest.NewRecorder()
		request, err := http.NewRequest(http.MethodDelete, "/speakers", reader)

		if err != nil {
			t.Fatal(err)
		}

		repositoryMock.On("GetById", mock.Anything).Return(&person.Person{}, nil)
		repositoryMock.On("Delete", mock.Anything).Return(nil)
		handler := http.HandlerFunc(controller.DeletePerson)

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
