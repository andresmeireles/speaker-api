package config

import (
	"encoding/json"
	"net/http"

	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
	web "github.com/andresmeireles/speaker/internal/web/decoder"
)

type ConfigController struct {
	configRepository ConfigRepository
	actions          Actions
}

func NewController(repo ConfigRepository, action Actions) ConfigController {
	return ConfigController{repo, action}
}

func (c ConfigController) New(s servicelocator.ServiceLocator) any {
	return ConfigController{
		configRepository: servicelocator.Get[Repository](s),
		actions:          servicelocator.Get[Actions](s),
	}
}

func (c ConfigController) WriteConfig(w http.ResponseWriter, r *http.Request) {
	body, err := web.DecodePostBody[[]Config](r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	for _, config := range body {
		err = c.actions.Write(config.Name, config.Value)
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Config created"))
}

func (c ConfigController) GetConfigs(w http.ResponseWriter, r *http.Request) {
	configs, err := c.configRepository.GetAll()
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
