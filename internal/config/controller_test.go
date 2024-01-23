package config_test

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/andresmeireles/speaker/internal/config"
// 	"github.com/andresmeireles/speaker/testdata"
// )

// func TestConfigController(t *testing.T) {
// 	controller := testdata.GetService[config.ConfigController]()

// 	t.Run("should return config", func(t *testing.T) {
// 		// arrange
// 		clearDB()
// 		repo := config.ConfigRepository{}
// 		err := repo.Add(config.Config{Name: "key", Value: "value"})
// 		_ = repo.Add(config.Config{Name: "key2", Value: "value"})

// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		req, err := http.NewRequest(http.MethodGet, "/configs", nil)

// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		rr := httptest.NewRecorder()
// 		handler := http.HandlerFunc(controller.GetConfigs)
// 		handler.ServeHTTP(rr, req)

// 		if status := rr.Code; status != http.StatusOK {
// 			t.Errorf("handler returned wrong status code: got %v want %v",
// 				status, http.StatusOK)
// 		}

// 		expect := "[{\"name\":\"key\",\"value\":\"value\"},{\"name\":\"key2\",\"value\":\"value\"}]"

// 		if rr.Body.String() != expect {
// 			t.Errorf("handler returned unexpected body: got %v want %v",
// 				rr.Body.String(), expect)
// 		}
// 	})
// }
