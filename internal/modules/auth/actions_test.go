package auth_test

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/andresmeireles/speaker/internal/modules/auth"
	"github.com/golang-jwt/jwt/v5"
)

func TestMain(m *testing.M) {
	randomNumber := rand.Intn(100000)
	hash := sha256.New()
	hash.Write([]byte(strconv.Itoa(randomNumber)))
	hashedBytes := hash.Sum(nil)
	h := hex.EncodeToString(hashedBytes)
	os.Setenv("APP_KEY", h)
	os.Exit(m.Run())
}

func createJwt(expireDate int64) string {
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "andres.meireles",
		"sub": "andres.meireles",
		"exp": expireDate,
	})

	tokenString, _ := tk.SignedString([]byte(os.Getenv("APP_KEY")))

	return tokenString
}

func TestValidateJson(t *testing.T) {
	t.Run("should return true", func(t *testing.T) {
		// arrange
		token := createJwt(time.Now().Add(time.Hour * 24).Unix())

		// act
		validateToken := auth.ValidateJwt(token)

		// assert
		if !validateToken {
			t.Fatalf("expected true, got false")
		}
	})

	t.Run("should return false when exp is invalid", func(t *testing.T) {
		// arrange
		token := createJwt(0)

		// act
		validateToken := auth.ValidateJwt(token)

		// assert
		if validateToken {
			t.Fatalf("expected false, got true")
		}
	})

	t.Run("should return false when token is invalid", func(t *testing.T) {
		// act
		validateToken := auth.ValidateJwt("jarl")

		// assert
		if validateToken {
			t.Fatalf("expected false, got true")
		}
	})
}
