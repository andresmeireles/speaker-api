package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

func Root() (string, error) {
	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	for {
		goModPath := filepath.Join(dir, "go.mod")
		_, err := os.Stat(goModPath)
		if err == nil {
			return dir, nil
		}
		if !os.IsNotExist(err) {
			return "", err
		}

		// go upwards
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}

	return "", fmt.Errorf("go.mod not found")
}
