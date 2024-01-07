package codesender

import (
	"fmt"
	"log/slog"
	"math/rand"
	"strconv"
	"time"

	"github.com/andresmeireles/speaker/internal/db/entity"
)

type Actions struct {
	repository AuthCodeRepository
}

const EXPIRE_TIME_MINUTES = 5

func (a Actions) CreateCode(user entity.User) (string, error) {
	randGenerator := rand.New(rand.NewSource(time.Now().Unix()))
	code := strconv.Itoa(randGenerator.Int())
	authCode := entity.AuthCode{
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

	if row.ExpiresAt.Before(time.Now()) {
		slog.Error("Auth code expired", "code", code)

		return fmt.Errorf("auth code expired")
	}

	if row.User.Email != userEmail {
		slog.Error("Invalid user", "email", userEmail)

		return fmt.Errorf("invalid user")
	}

	return nil
}
