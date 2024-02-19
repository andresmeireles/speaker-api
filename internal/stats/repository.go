package stats

import (
	"github.com/andresmeireles/speaker/internal/person"
	"github.com/andresmeireles/speaker/internal/repository"
)

type StatsRepository interface {
	GetNumberOfSpeakersForPersons() ([]SpeakerReport, error)
}

type Repository struct {
	repository       repository.RepositoryInterface
	personRepository person.PersonRepository
}

func NewRepository(
	repository repository.RepositoryInterface,
	personRepository person.PersonRepository,
) StatsRepository {
	return Repository{
		repository:       repository,
		personRepository: personRepository,
	}
}

func (r Repository) GetNumberOfSpeakersForPersons() ([]SpeakerReport, error) {
	query := "select p.id as speaker_name, " +
	"(select count(i.id) from invites i where i.status = 5 and i.person_id = p.id) as done_speaks, " +
	"(select count(i2.id) from invites i2 where i2.person_id = p.id ) as total_speaks " +
	"from persons p"

	rows, err := r.repository.Query(query)

	if err != nil {
		return nil, err
	}

	speakerReports := make([]SpeakerReport, 0)

	for rows.Next() {
		sr := new(SpeakerReport)
		if err := rows.Scan(&sr.speaker_id, &sr.doneSpeaks, &sr.totalSpeaks); err != nil {
			return nil, err
		}

		person, err := r.personRepository.GetById(sr.speaker_id)
		if err != nil {
			return nil, err
		}

		sr.speaker = *person
		speakerReports = append(speakerReports, *sr)
	}

	return speakerReports, nil
}
