package stats

import (
	"encoding/json"
	"net/http"

	"github.com/andresmeireles/speaker/internal/tools/responses"
)

type StatsController struct {
	repository StatsRepository
}

func NewStatsController(repository StatsRepository) StatsController {
	return StatsController{
		repository: repository,
	}
}

func (c StatsController) SpeakersStats(w http.ResponseWriter, r *http.Request) {
	report, err := c.repository.GetNumberOfSpeakersForPersons()
	if err != nil {
		responses.BadResponse(w, err)

		return
	}

	response := make([]speakerReport, 0)

	for _, speaker := range report {
		response = append(response, ExportSpeakerReport(speaker))
	}

	res, _ := json.Marshal(response)

	responses.Ok(w, res)
}
