package stats

import "github.com/andresmeireles/speaker/internal/person"

type speakerReport struct {
	Speaker person.Person `json:"speaker"`
	Speaks  int           `json:"speaks"`
}

func ExportSpeakerReport(report SpeakerReport) speakerReport {
	return speakerReport{
		Speaker: report.Speaker(),
		Speaks:  report.Speaks(),
	}
}
