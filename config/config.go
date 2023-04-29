package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Configuration struct {
	ENV       string
	SentryDSN string
	SecretKey string
	Database  DatabaseConfiguration
	Server    ServerConfiguration
	SMTP      SMTPConfiguration
	Storage   StorageConfiguration
}

type ServerConfiguration struct {
	Domain   string
	Port     string
	Timeout  int
	MediaUrl string
}

type DatabaseConfiguration struct {
	Name     string
	Username string
	Password string
	Host     string
	Port     string
}

type SMTPConfiguration struct {
	From     string
	Host     string
	Port     string
	Password string
}

type StorageConfiguration struct {
	MediaRoot string
}

func Init() *Configuration {
	err := godotenv.Load("./config/.env")
	if err != nil {
		if _, err := os.Stat("./config/.env"); os.IsNotExist(err) {
			os.Create("./config/.env")
		}
	}

	return &Configuration{
		ENV:       os.Getenv("ENV"),
		SecretKey: os.Getenv("SECRET_KEY"),
		SentryDSN: os.Getenv("SENTRY_DSN"),
		Database: DatabaseConfiguration{
			Name:     os.Getenv("DB_NAME"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
		},
		Server: ServerConfiguration{
			Domain:  os.Getenv("SW_DOMAIN"),
			Port:    os.Getenv("SW_PORT"),
			Timeout: convertInt("SW_TIMEOUT"),
		},
		SMTP: SMTPConfiguration{
			From:     os.Getenv("SMTP_FROM"),
			Host:     os.Getenv("SMTP_HOST"),
			Port:     os.Getenv("SMTP_PORT"),
			Password: os.Getenv("SMTP_PASSWORD"),
		},
	}
}

func convertInt(envKey string) int {
	value, err := strconv.Atoi(os.Getenv(envKey))
	if err != nil {
		return 0
	}
	return value
}
