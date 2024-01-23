package auth_test

// import (
// 	"os"
// 	"testing"
// 	"time"

// 	"github.com/andresmeireles/speaker/internal/auth"
// 	"github.com/andresmeireles/speaker/testdata"
// 	"github.com/golang-jwt/jwt/v5"
// )

// func TestMain(m *testing.M) {
// 	testdata.SetupDatabase(m)
// }

// func createJwt(expireDate int64) string {
// 	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"iss": "andres.meireles@email.com",
// 		"sub": "andres.meireles@email.com",
// 		"exp": expireDate,
// 	})

// 	tokenString, _ := tk.SignedString([]byte(os.Getenv("APP_KEY")))

// 	return tokenString
// }

// func TestValidateJson(t *testing.T) {
// 	actions := testdata.GetService[auth.Actions]()

// 	t.Run("should return true", func(t *testing.T) {
// 		// arrange
// 		token := createJwt(time.Now().Add(time.Hour * 24).Unix())

// 		// act
// 		validateToken := actions.ValidateJwt(token)

// 		// assert
// 		if !validateToken {
// 			t.Fatalf("expected true, got false")
// 		}
// 	})

// 	t.Run("should return false when exp is invalid", func(t *testing.T) {
// 		// arrange
// 		token := createJwt(0)

// 		// act
// 		validateToken := actions.ValidateJwt(token)

// 		// assert
// 		if validateToken {
// 			t.Fatalf("expected false, got true")
// 		}
// 	})

// 	t.Run("should return false when token is invalid", func(t *testing.T) {
// 		// act
// 		validateToken := actions.ValidateJwt("jarl")

// 		// assert
// 		if validateToken {
// 			t.Fatalf("expected false, got true")
// 		}
// 	})
// }
