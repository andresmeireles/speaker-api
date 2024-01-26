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
	query := "SELECT p.id as speaker_id, COUNT(i.id) as speaks " +
		"FROM invites i JOIN persons p ON p.id = i.person_id GROUP BY p.id"
	rows, err := r.repository.Query(query)

	if err != nil {
		return nil, err
	}

	speakerReports := make([]SpeakerReport, 0)

	for rows.Next() {
		sr := new(SpeakerReport)
		if err := rows.Scan(&sr.speaker_id, &sr.speaks); err != nil {
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
