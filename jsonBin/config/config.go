package config

import (
	"errors"
	"github.com/joho/godotenv"
	"jsonBin/print"
	"os"
)

type Config struct {
	Key       string
	AccessKey string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		print.Error(errors.New("no .env file found"))
	}

	return &Config{
		Key:       os.Getenv("KEY"),
		AccessKey: os.Getenv("ACCESS"),
	}
}
