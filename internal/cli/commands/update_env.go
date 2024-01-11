package commands

import (
	"fmt"
	"os"
	"strings"
)

func modifyEnvFile(key, value string) error {
	root, err := Root()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fileBytes, err := os.ReadFile(root + "/.env")
	if err != nil {
		return fmt.Errorf("failed to read .env file: %w", err)
	}

	fileContent := string(fileBytes)

	lines := strings.Split(fileContent, "\n")
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), key+"=") {
			lines[i] = key + "=" + value

			break
		}
	}

	updatedContent := strings.Join(lines, "\n")

	err = os.WriteFile(".env", []byte(updatedContent), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write updated .env file: %w", err)
	}

	return nil
}
