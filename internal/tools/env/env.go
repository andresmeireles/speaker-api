package env

import (
	"fmt"
	"os"
	"strings"
)

func IsDev() bool {
	return os.Getenv("APP_MODE") == "dev"
}

func AppKey() (string, error) {
	key := os.Getenv("APP_KEY")
	if len(strings.Trim(key, " ")) == 0 {
		return "", fmt.Errorf("APP_KEY not set")
	}

	return key, nil
}

func ShowStackTrace() bool {
	return os.Getenv("SHOW_STACK_TRACE") == "true"
}

func ShowErrorFile() bool {
	return os.Getenv("SHOW_ERROR_FILE") == "true"
}
