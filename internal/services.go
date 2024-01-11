package internal

import (
	"github.com/andresmeireles/speaker/internal/db"
	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/db/repository"
	"github.com/andresmeireles/speaker/internal/modules/auth"
	"github.com/andresmeireles/speaker/internal/modules/codesender"
	"github.com/andresmeireles/speaker/internal/modules/config"
	"github.com/andresmeireles/speaker/internal/modules/invite"
	"github.com/andresmeireles/speaker/internal/modules/person"
	"github.com/andresmeireles/speaker/internal/modules/user"
	"github.com/andresmeireles/speaker/internal/tools"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
)

func controllers() []servicelocator.Instantiable {
	return []servicelocator.Instantiable{
		user.UserController{},
		config.ConfigController{},
		auth.AuthController{},
		invite.InviteController{},
		person.PersonController{},
	}
}

func repos() []servicelocator.Instantiable {
	return []servicelocator.Instantiable{
		user.UserRepository{},
		config.ConfigRepository{},
		codesender.AuthCodeRepository{},
		auth.AuthRepository{},
		person.PersonRepository{},
		invite.InviteRepository{},
	}
}

func actions() []servicelocator.Instantiable {
	return []servicelocator.Instantiable{
		config.Actions{},
		codesender.Actions{},
		auth.Actions{},
		invite.Actions{},
		person.Actions{},
	}
}

func misc() []servicelocator.Instantiable {
	return []servicelocator.Instantiable{
		db.Connection{},

		&tools.Email{},

		repository.Repository[entity.Auth]{},
		repository.Repository[entity.AuthCode]{},
		repository.Repository[entity.Person]{},
		repository.Repository[entity.Config]{},
		repository.Repository[entity.User]{},
		repository.Repository[entity.Invite]{},
	}
}

func Services() []servicelocator.Instantiable {
	services := make([]servicelocator.Instantiable, 0)
	services = append(services, misc()...)
	services = append(services, repos()...)
	services = append(services, actions()...)
	services = append(services, controllers()...)
	// return []servicelocator.Instantiable{
	// 	db.Connection{},

	// 	repository.Repository[entity.Auth]{},
	// 	repository.Repository[entity.AuthCode]{},
	// 	repository.Repository[entity.Person]{},
	// 	repository.Repository[entity.Config]{},
	// 	repository.Repository[entity.User]{},
	// 	repository.Repository[entity.Invite]{},

	// 	user.UserRepository{},
	// 	user.UserController{},

	// 	config.ConfigRepository{},
	// 	config.Actions{},
	// 	config.ConfigController{},

	// 	codesender.Actions{},
	// 	codesender.AuthCodeRepository{},

	// 	auth.Actions{},
	// 	auth.AuthRepository{},
	// 	auth.AuthController{},

	// 	invite.Actions{},
	// 	invite.InviteRepository{},
	// 	invite.InviteController{},

	// 	person.Actions{},
	// 	person.PersonRepository{},
	// 	person.PersonRepository{},
	// }
	return services
}
