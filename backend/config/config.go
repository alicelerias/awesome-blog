package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host       string
	Driver     string
	Database   string
	Port       string
	User       string
	Password   string
	LogLevel   string
	JWTSecret  []byte
	AuthCookie string
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		godotenv.Load()

		config = &Config{
			Host:       os.Getenv("DB_HOST"),
			Driver:     os.Getenv("DB_DRIVER"),
			Database:   os.Getenv("DB_NAME"),
			Port:       os.Getenv("DB_PORT"),
			User:       os.Getenv("DB_USER"),
			Password:   os.Getenv("DB_PASSWORD"),
			LogLevel:   os.Getenv("LOG_LEVEL"),
			JWTSecret:  []byte(os.Getenv("JWT_SECRET")),
			AuthCookie: os.Getenv("AUTH_COOKIE"),
		}
	}

	return config

}
