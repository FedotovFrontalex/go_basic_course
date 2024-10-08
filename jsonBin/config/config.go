package config

import (
	"errors"
	"jsonBin/print"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		print.Error(errors.New("no .env file found"))
	}

	return &Config{
		Key: os.Getenv("KEY"),
	}
}
