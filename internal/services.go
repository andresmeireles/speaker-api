package internal

import (
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
		codesender.AuthCodeRepository{},
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

		repository.Repository[auth.Auth]{},
		repository.Repository[codesender.AuthCode]{},
		repository.Repository[person.Person]{},
		repository.Repository[config.Config]{},
		repository.Repository[user.User]{},
		repository.Repository[invite.Invite]{},
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
