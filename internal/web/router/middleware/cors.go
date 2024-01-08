package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/cors"
)

func origins() []string {
	origins := os.Getenv("CORS_ALLOWED_ORIGINS")
	if origins == "" {
		return []string{"*"}
	}

	if !strings.Contains(origins, ",") {
		return []string{origins}
	}

	return strings.Split(origins, ",")
}

func Cors(next http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins:   origins(),
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})(next)
}
