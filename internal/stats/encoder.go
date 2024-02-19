package stats

import "github.com/andresmeireles/speaker/internal/person"

type speakerReport struct {
	Speaker person.Person `json:"speaker"`
	DoneSpeaks int           `json:"doneSpeaks"`
	TotalSpeaks int       `json:"totalSpeaks"`
}

func ExportSpeakerReport(report SpeakerReport) speakerReport {
	return speakerReport{
		Speaker: report.Speaker(),
		DoneSpeaks:  report.DoneSpeaks(),
		TotalSpeaks: report.TotalSpeaks(),
	}
}
