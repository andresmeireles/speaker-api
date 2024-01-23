package stats

import "github.com/andresmeireles/speaker/internal/person"

type SpeakerReport struct {
	speaker    person.Person
	speaker_id int
	speaks     int
}
