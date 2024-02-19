package stats

import (
	"github.com/andresmeireles/speaker/internal/person"
)

type SpeakerReport struct {
	speaker    person.Person
	speaker_id int
	doneSpeaks     int
	totalSpeaks int
}

func (sr SpeakerReport) Speaker() person.Person {
	return sr.speaker
}

func (sr SpeakerReport) DoneSpeaks() int {
	return sr.doneSpeaks
}

func (sr SpeakerReport) TotalSpeaks() int {
	return sr.totalSpeaks
}
