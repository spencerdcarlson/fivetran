package api

import (
	"errors"
	"fmt"
	"os"
)

type Auth struct {
	APISecret string
	APIKey    string
}

func BuildAuth(key string) (*Auth, error) {
	secret := os.Getenv("API_SECRET")
	if secret == "" {
		return nil, errors.New("API_SECRET environment variable not set")
	}
	envKey := os.Getenv("API_KEY")
	if key != "" {
		return &Auth{
			APISecret: secret,
			APIKey:    key,
		}, nil
	} else if envKey != "" {
		return &Auth{
			APISecret: secret,
			APIKey:    envKey,
		}, nil
	} else {
		return nil, errors.New(fmt.Sprintf("invalid API Key"))
	}
}
