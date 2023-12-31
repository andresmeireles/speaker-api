package auth

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/andresmeireles/speaker/internal/modules/codesender"
	"github.com/andresmeireles/speaker/internal/modules/user"
	"github.com/andresmeireles/speaker/internal/tools"
	"github.com/golang-jwt/jwt/v5"
)

const HOURS_TO_EXPIRE = 24

type Actions struct {
	repository       AuthRepository
	userRepository   user.UserRepository
	email            *tools.Email
	codeSenderAction codesender.Actions
}

func (a Actions) Logout(userId int) error {
	err := a.repository.ExpireTokenByUserId(userId)
	if err != nil {
		slog.Error("error on update auths to expire", err)
	}

	return err
}

func (a Actions) ValidateJwt(token string) bool {
	parseFunc := func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(os.Getenv("APP_KEY")), nil
	}
	jwtToken, err := jwt.Parse(token, parseFunc)

	if err != nil {
		return false
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		slog.Error("exp claims are not float64")

		return false
	}

	return int64(exp) > time.Now().Unix()
}

func (a Actions) CreateJWT(user entity.User) (entity.Auth, error) {
	appKey := os.Getenv("APP_KEY")
	if appKey == "" {
		return entity.Auth{}, fmt.Errorf("APP_KEY not set")
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "andres.meireles",
		"sub": user.Email,
		"exp": time.Now().Add(time.Hour * HOURS_TO_EXPIRE).Unix(),
	})
	token, err := jwtToken.SignedString([]byte(appKey))

	if err != nil {
		return entity.Auth{}, err
	}

	newAuth := entity.Auth{
		User:   user,
		UserId: user.Id,
		Hash:   token,
	}
	err = a.repository.Add(newAuth)

	if err != nil {
		return entity.Auth{}, err
	}

	return newAuth, nil
}

func (a Actions) HasEmail(email string) bool {
	_, err := a.userRepository.GetByEmail(email)

	return err == nil
}

func (a Actions) SendCode(email string) error {
	user, err := a.userRepository.GetByEmail(email)
	if err != nil {
		return err
	}

	code, err := a.codeSenderAction.CreateCode(user)
	if err != nil {
		return err
	}

	err = a.email.Send(code, user.Email)

	return err
}
