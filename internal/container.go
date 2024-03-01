package internal

import (
	"os"

	"github.com/andresmeireles/di"
	"github.com/andresmeireles/speaker/internal/auth"
	"github.com/andresmeireles/speaker/internal/cli/auxcmd"
	"github.com/andresmeireles/speaker/internal/codesender"
	"github.com/andresmeireles/speaker/internal/config"
	"github.com/andresmeireles/speaker/internal/db"
	"github.com/andresmeireles/speaker/internal/invite"
	"github.com/andresmeireles/speaker/internal/person"
	"github.com/andresmeireles/speaker/internal/repository"
	"github.com/andresmeireles/speaker/internal/stats"
	"github.com/andresmeireles/speaker/internal/tools"
	"github.com/andresmeireles/speaker/internal/user"
	"github.com/andresmeireles/speaker/internal/web/router"
	"github.com/go-chi/chi/v5"
)

func DependenciesContainer() []di.Dependency {
	return []di.Dependency{
		di.NewTypedDependency[db.Connection](db.NewConnection),
		di.NewTypedDependency[tools.E](func() *tools.Email {
			host := os.Getenv("SMTP_HOST")
			port := os.Getenv("SMTP_PORT")
			password := os.Getenv("SMTP_PASSWORD")
			email := os.Getenv("SMTP_USER")

			return tools.NewEmail(host, password, port, email)
		}),
		di.NewTypedDependency[*tools.Email](func() *tools.Email {
			host := os.Getenv("SMTP_HOST")
			port := os.Getenv("SMTP_PORT")
			password := os.Getenv("SMTP_PASSWORD")
			email := os.Getenv("SMTP_USER")

			return tools.NewEmail(host, password, port, email)
		}),

		// repository
		di.NewTypedDependency[repository.RepositoryInterface](repository.NewRepository),
		di.NewTypedDependency[repository.Repository](repository.NewRepository),
		di.NewTypedDependency[codesender.Repository](codesender.NewRepository),
		di.NewTypedDependency[codesender.CodeSenderRepository](codesender.NewRepository),
		di.NewTypedDependency[person.PersonRepository](person.NewRepository),
		di.NewTypedDependency[person.Repository](person.NewRepository),
		di.NewTypedDependency[user.UserRepository](user.NewRepository),
		di.NewTypedDependency[user.Repository](user.NewRepository),
		di.NewTypedDependency[invite.InviteRepository](invite.NewRepository),
		di.NewTypedDependency[auth.AuthRepository](auth.NewRepository),
		di.NewTypedDependency[config.ConfigRepository](config.NewRepository),
		di.NewTypedDependency[stats.StatsRepository](stats.NewRepository),

		// actions / service
		di.NewTypedDependency[codesender.Service](codesender.NewAction),
		di.NewTypedDependency[auth.Actions](auth.NewAction),
		di.NewTypedDependency[config.Actions](config.NewActions),
		di.NewTypedDependency[person.ActionsInterface](person.NewAction),
		di.NewTypedDependency[person.Actions](person.NewAction),
		di.NewTypedDependency[invite.InviteService](invite.NewAction),

		// controller
		di.NewTypedDependency[person.PersonController](person.NewController),
		di.NewTypedDependency[user.UserController](user.NewController),
		di.NewTypedDependency[config.ConfigController](config.NewController),
		di.NewTypedDependency[auth.AuthController](auth.NewController),
		di.NewTypedDependency[invite.InviteController](invite.NewController),
		di.NewTypedDependency[stats.StatsController](stats.NewStatsController),

		// cli cmd
		di.NewTypedDependency[auxcmd.Migration](auxcmd.NewMigration),

		// server
		di.NewTypedDependency[*chi.Mux](func() *chi.Mux {
			return chi.NewRouter()
		}),
		di.NewTypedDependency[router.Router](router.NewRouter),
	}
}
