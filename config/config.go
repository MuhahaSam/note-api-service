package config

import (
	"fmt"
	"os"
)

type DbConfig struct {
	Engine string
	Dsn    string
}

type Config struct {
	DbConfig DbConfig
}

var config *Config = nil

func GetConfig() *Config {
	if config != nil {
		return config
	}

	config = &Config{
		DbConfig: DbConfig{
			Engine: os.Getenv("DB_ENGINE"),
			Dsn: fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
				os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_NAME"),
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASSWORD"),
			),
		},
	}

	return config
}
