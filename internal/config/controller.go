package config

import (
	"encoding/json"
	"net/http"

	"github.com/andresmeireles/speaker/internal/tools/responses"
	web "github.com/andresmeireles/speaker/internal/web/decoder"
)

type ConfigController struct {
	configRepository ConfigRepository
	actions          Actions
}

func NewController(repo ConfigRepository, action Actions) ConfigController {
	return ConfigController{repo, action}
}

func (c ConfigController) WriteConfig(w http.ResponseWriter, r *http.Request) {
	body, err := web.DecodePostBody[[]Config](r.Body)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	for _, config := range body {
		err = c.actions.Write(config.Name, config.Value)
	}

	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	responses.Created(w, []byte("Configs successfully created"))
}

func (c ConfigController) GetConfigs(w http.ResponseWriter, r *http.Request) {
	configs, err := c.configRepository.GetAll()
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	jsonConfigs, err := json.Marshal(configs)
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	responses.Ok(w, jsonConfigs)
}
