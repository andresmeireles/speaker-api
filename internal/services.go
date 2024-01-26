package internal

import (
	"os"

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
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
	"github.com/andresmeireles/speaker/internal/user"
)

func DIContainer(sl *servicelocator.SL) {
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
		servicelocator.AddDependency[repository.RepositoryInterface](repository.NewRepository),
		servicelocator.AddDependency[repository.Repository](repository.NewRepository),
		servicelocator.AddDependency[codesender.CodeSenderRepository](codesender.NewRepository),
		servicelocator.AddDependency[person.PersonRepository](person.NewRepository),
		servicelocator.AddDependency[user.Repository](user.NewRepository),
		servicelocator.AddDependency[user.UserRepository](user.NewRepository),
		servicelocator.AddDependency[invite.InviteRepository](invite.NewRepository),
		servicelocator.AddDependency[person.Repository](person.NewRepository),
		servicelocator.AddDependency[auth.AuthRepository](auth.NewRepository),
		servicelocator.AddDependency[config.ConfigRepository](config.NewRepository),
		servicelocator.AddDependency[stats.StatsRepository](stats.NewRepository),

		// action / service
		servicelocator.AddDependency[codesender.Actions](codesender.NewAction),
		servicelocator.AddDependency[auth.Actions](auth.NewAction),
		servicelocator.AddDependency[config.Actions](config.NewActions),
		servicelocator.AddDependency[person.ActionsInterface](person.NewAction),
		servicelocator.AddDependency[invite.InviteService](invite.NewAction),

		// controller
		servicelocator.AddDependency[person.PersonController](person.NewController),
		servicelocator.AddDependency[user.UserController](user.NewController),
		servicelocator.AddDependency[config.ConfigController](config.NewController),
		servicelocator.AddDependency[auth.AuthController](auth.NewController),
		servicelocator.AddDependency[invite.InviteController](invite.NewController),
		servicelocator.AddDependency[stats.StatsController](stats.NewStatsController),

		// cli cmd
		servicelocator.AddDependency[auxcmd.Migration](auxcmd.NewMigration),
	}

	servicelocator.Set(sl, injections)
}
