package config

import "os"

type Config struct {
	Database Database
}

type Database struct {
	Host     string
	User     string
	Port     string
	Password string
	Database string
}

func GetConfig() Config {
	return Config{
		Database: Database{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Port:     os.Getenv("DB_PORT"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_NAME"),
		},
	}
}
