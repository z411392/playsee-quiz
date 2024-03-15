package auth

import (
	"errors"
	"os"
)

func ValidateApiKey(apiKey string) error {
	if apiKey != os.Getenv("API_KEY") {
		return errors.New("invalid Key")
	}
	return nil
}
