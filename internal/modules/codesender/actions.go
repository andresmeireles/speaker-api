package codesender

import (
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
		ExpiresAt: int(time.Now().Add(EXPIRE_TIME_MINUTES * time.Minute).Unix()),
	}
	err := a.repository.Add(authCode)

	if err != nil {
		return "", err
	}

	return code, nil
}
