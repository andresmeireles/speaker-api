package codesender

import (
	"fmt"
	"log/slog"
	"math/rand"
	"strconv"
	"time"

	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
	"github.com/andresmeireles/speaker/internal/user"
)

type Actions struct {
	repository repositoryInterface
}

const EXPIRE_TIME_MINUTES = 5

func NewAction(repository repositoryInterface) Actions {
	return Actions{
		repository: repository,
	}
}

func (a Actions) New(s servicelocator.ServiceLocator) any {
	return Actions{
		repository: servicelocator.Get[Repository](s),
	}
}

func (a Actions) CreateCode(user user.User) (string, error) {
	randGenerator := rand.New(rand.NewSource(time.Now().Unix()))
	code := strconv.Itoa(randGenerator.Int())
	authCode := AuthCode{
		UserId:    user.Id,
		User:      user,
		Code:      code,
		ExpiresAt: time.Now().Add(EXPIRE_TIME_MINUTES * time.Minute),
	}
	err := a.repository.Add(authCode)

	if err != nil {
		return "", err
	}

	return code, nil
}

func (a Actions) VerifyCode(userEmail, code string) error {
	row, err := a.repository.GetByCode(code)
	if err != nil {
		slog.Error("Failed to get auth code", err)

		return err
	}

	// Em main o timezone esta definido para America/SaoPaulo ou -3
	// mas o banco de dados salva o timestamp sem zona.
	//
	// Devo verficar como criar um now sem timezone, o que acho dificil.
	//
	// A solução mais facil foi de atrazar o horario em 3 horas, para que
	// quando for resolvido o horario e adicionar as 3 horas ao horario atual
	// por causa do UTC a validação ainda funcione.
	if row.ExpiresAt.Before(time.Now().Add(time.Hour * -3)) {
		slog.Error("Auth code expired", "current time", time.Now(), "expired time", row.ExpiresAt)

		return fmt.Errorf("auth code expired")
	}

	if row.User.Email != userEmail {
		slog.Error("Invalid user", "email", userEmail)

		return fmt.Errorf("invalid user")
	}

	return nil
}
