package config

import (
	"log"
	"os"
	"strconv"
)

type DbConfig struct {
	Engine   string
	Host     string
	Port     int
	DbName   string
	User     string
	Password string
}

type Config struct {
	DbConfig DbConfig
}

var config *Config = nil

func GetConfig() *Config {
	if config != nil {
		return config
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("db port is not number: %s", err)
	}

	config = &Config{
		DbConfig: DbConfig{
			Engine:   os.Getenv("DB_ENGINE"),
			Host:     os.Getenv("DB_HOST"),
			Port:     port,
			DbName:   os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
		},
	}

	return config
}
