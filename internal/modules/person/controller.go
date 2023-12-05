package person

import (
	"net/http"
	"os"
)

func ShowMode(w http.ResponseWriter, r *http.Request) {
	mode := os.Getenv("MODE")

	w.Write([]byte(mode))
}
