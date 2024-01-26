package stats

import (
	"github.com/andresmeireles/speaker/internal/person"
)

type SpeakerReport struct {
	speaker    person.Person
	speaker_id int
	speaks     int
}

func (sr SpeakerReport) Speaker() person.Person {
	return sr.speaker
}

func (sr SpeakerReport) Speaks() int {
	return sr.speaks
}
