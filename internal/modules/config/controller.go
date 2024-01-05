package config

import (
	"encoding/json"
	"net/http"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/web"
)

func WriteConfig(w http.ResponseWriter, r *http.Request) {
	body, err := web.DecodePostBody[[]entity.Config](r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	for _, config := range body {
		err = Write(config.Name, config.Value, ConfigRepository{})
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Config created"))
}

func GetConfigs(w http.ResponseWriter, r *http.Request) {
	repository := ConfigRepository{}
	configs, err := repository.GetAll()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error when get configs"))

		return
	}

	jsonConfigs, err := json.Marshal(configs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error when get configs"))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonConfigs))
}
