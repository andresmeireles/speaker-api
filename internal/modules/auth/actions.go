package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/andresmeireles/speaker/internal/db/entity"
	"github.com/golang-jwt/jwt/v5"
)

func ExpireAuth(auth entity.Auth, repository AuthRepository) error {
	auth.Expired = true
	return repository.Update(auth)
}

func ValidateJwt(token string) bool {
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
	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Unix() {
		return false
	}

	return true
}

func CreateJWT(user entity.User, repository AuthRepository) (entity.Auth, error) {
	appKey := os.Getenv("APP_KEY")

	if appKey == "" {
		return entity.Auth{}, fmt.Errorf("APP_KEY not set")
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "andres.meireles",
		"sub": user.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := jwtToken.SignedString([]byte(appKey))

	if err != nil {
		return entity.Auth{}, err
	}

	newAuth := entity.Auth{
		User: user,
		Hash: token,
	}
	err = repository.Add(newAuth)
	if err != nil {
		return entity.Auth{}, err
	}

	return newAuth, nil
}

func CheckCode(userId int, token string, repository AuthRepository) error {
	code, err := repository.AuthCodeByUser(token, userId)
	if err != nil {
		return err
	}
	if code == nil {
		return fmt.Errorf("No code auth found")
	}
	if int64(code.ExpiresAt) < time.Now().Unix() {
		return fmt.Errorf("Code is expired")
	}

	return nil
}
