package stats

import "github.com/andresmeireles/speaker/internal/repository"

type Repository struct {
	repository repository.RepositoryInterface
}

func NewRepository(repository repository.RepositoryInterface) Repository {
	return Repository{
		repository: repository,
	}
}

func (r Repository) GetNumberOfSpeakersForPersons() {

}
