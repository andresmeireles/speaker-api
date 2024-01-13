package internal

import (
	"os"

	"github.com/andresmeireles/speaker/internal/auth"
	"github.com/andresmeireles/speaker/internal/codesender"
	"github.com/andresmeireles/speaker/internal/config"
	"github.com/andresmeireles/speaker/internal/db"
	"github.com/andresmeireles/speaker/internal/invite"
	"github.com/andresmeireles/speaker/internal/person"
	"github.com/andresmeireles/speaker/internal/repository"
	"github.com/andresmeireles/speaker/internal/tools"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
	"github.com/andresmeireles/speaker/internal/user"
)

func Controllers() []servicelocator.Instantiable {
	return []servicelocator.Instantiable{
		user.UserController{},
		config.ConfigController{},
		auth.AuthController{},
		invite.InviteController{},
		person.PersonController{},
	}
}

func Repos() []servicelocator.Instantiable {
	return []servicelocator.Instantiable{
		user.UserRepository{},
		config.ConfigRepository{},
		codesender.Repository{},
		auth.AuthRepository{},
		person.PersonRepository{},
		invite.InviteRepository{},
	}
}

func Actions() []servicelocator.Instantiable {
	return []servicelocator.Instantiable{
		config.Actions{},
		codesender.Actions{},
		auth.Actions{},
		invite.Actions{},
		person.Actions{},
	}
}

func Misc() []servicelocator.Instantiable {
	return []servicelocator.Instantiable{
		db.Connection{},

		&tools.Email{},

		repository.Repository{},
	}
}

func Services() []servicelocator.Instantiable {
	services := make([]servicelocator.Instantiable, 0)
	services = append(services, Misc()...)
	services = append(services, Repos()...)
	services = append(services, Actions()...)
	services = append(services, Controllers()...)

	return services
}

func DIContainer(sl *servicelocator.ServiceLocator) {
	injections := []servicelocator.Dependency{
		// with no deps
		servicelocator.AddDependency[db.Connection](db.NewConnection),
		servicelocator.AddDependency[*tools.Email](func() *tools.Email {
			host := os.Getenv("SMTP_HOST")
			port := os.Getenv("SMTP_PORT")
			password := os.Getenv("SMTP_PASSWORD")
			email := os.Getenv("SMTP_USER")

			return tools.NewEmail(host, password, port, email)
		}),

		// repository
		servicelocator.AddInterface("repository.RepositoryInterface", repository.NewRepository),
		servicelocator.AddDependency[repository.Repository](repository.NewRepository),
		servicelocator.AddInterface("codesender.repositoryInterface", codesender.NewRepository),
		servicelocator.AddInterface("person.PersonRepositoryInterface", person.NewRepository),
		servicelocator.AddDependency[user.UserRepository](user.NewRepository),
		servicelocator.AddDependency[invite.InviteRepository](invite.NewRepository),
		servicelocator.AddDependency[person.PersonRepository](person.NewRepository),
		servicelocator.AddDependency[auth.AuthRepository](auth.NewRepository),
		servicelocator.AddDependency[config.ConfigRepository](config.NewRepository),

		// action
		servicelocator.AddDependency[codesender.Actions](codesender.NewAction),
		servicelocator.AddDependency[invite.Actions](invite.NewAction),
		servicelocator.AddDependency[auth.Actions](auth.NewAction),
		servicelocator.AddDependency[config.Actions](config.NewActions),
		servicelocator.AddInterface("person.ActionsInterface", person.NewAction),

		// controller
		servicelocator.AddDependency[person.PersonController](person.NewController),
		servicelocator.AddDependency[user.UserController](user.NewController),
		servicelocator.AddDependency[config.ConfigController](config.NewController),
		servicelocator.AddDependency[auth.AuthController](auth.NewController),
		servicelocator.AddDependency[invite.InviteController](invite.NewController),
	}

	servicelocator.SetRecursive(sl, injections)
}
