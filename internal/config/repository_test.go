package config_test

// import (
// 	"testing"

// 	"github.com/andresmeireles/speaker/internal/config"
// 	"github.com/andresmeireles/speaker/internal/repository"
// 	"github.com/andresmeireles/speaker/testdata"
// )

// func clearDB() {
// 	repository := testdata.GetService[repository.Repository]()
// 	q := "DELETE FROM configs"
// 	repository.SingleQuery(q)
// }

// func TestCreateConfig(t *testing.T) {
// 	clearDB()

// 	// arrange
// 	conf := config.Config{
// 		Name:  "key",
// 		Value: "value",
// 	}
// 	conf2 := config.Config{
// 		Name:  "key2",
// 		Value: "value2",
// 	}
// 	repository := testdata.GetService[config.ConfigRepository]()
// 	err := repository.Add(conf)

// 	if err != nil {
// 		t.Fatalf("expected nil, got %s", err)
// 	}

// 	err = repository.Add(conf2)

// 	if err != nil {
// 		t.Fatalf("expected nil, got %s", err)
// 	}

// 	// act
// 	configs, err := repository.GetAll()

// 	// assert
// 	if err != nil {
// 		t.Fatalf("expected nil, got %s", err)
// 	}

// 	if len(configs) != 2 {
// 		t.Fatalf("expected 2, got %d", len(configs))
// 	}
// }

// func TestFailToCreateWithSameId(t *testing.T) {
// 	// arrange
// 	conf := config.Config{
// 		Name:  "key3",
// 		Value: "value",
// 	}
// 	repository := testdata.GetService[config.ConfigRepository]()
// 	err := repository.Add(conf)

// 	if err != nil {
// 		t.Fatalf("expected nil, got %s", err)
// 	}

// 	conf2 := config.Config{
// 		Name:  "key3",
// 		Value: "value2",
// 	}

// 	// act
// 	err = repository.Add(conf2)

// 	// assert
// 	if err == nil {
// 		t.Fatalf("expected error, got nil")
// 	}

// 	if err.Error() != "config with name key3 already exists" {
// 		t.Fatalf("expected config with name key already exists, got %s", err.Error())
// 	}
// }

// func TestReturnById(t *testing.T) {
// 	clearDB()

// 	// arrange
// 	conf := config.Config{
// 		Name:  "key4",
// 		Value: "value",
// 	}
// 	repository := testdata.GetService[config.ConfigRepository]()
// 	err := repository.Add(conf)

// 	if err != nil {
// 		t.Fatalf("expected nil, got %s", err)
// 	}

// 	c, err := repository.GetByName("key4")
// 	if err != nil {
// 		t.Fatalf("expected nil, got %s", err)
// 	}

// 	// act
// 	config, err := repository.GetById(c.Id)

// 	// assert
// 	if err != nil {
// 		t.Fatalf("expected nil, got %s", err)
// 	}

// 	if config.Name != "key4" {
// 		t.Fatalf("expected key4, got %s", config.Name)
// 	}
// }
