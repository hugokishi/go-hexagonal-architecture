package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

func ReadConfig(filename string) error {
	if os.Getenv("APP_ENV") == "" || os.Getenv("APP_ENV") == "local" {
		if err := godotenv.Load(filename); err != nil {
			return errors.Wrap(err, "")
		}
	}
	return nil
}
