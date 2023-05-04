package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Limit        string
	Host         string
	Driver       string
	Database     string
	Port         string
	User         string
	Password     string
	LogLevel     string
	JWTSecret    []byte
	AuthCookie   string
	AllowedHosts string
	RedisPort    string
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		godotenv.Load()

		config = &Config{
			Limit:        os.Getenv("LIMIT"),
			Host:         os.Getenv("DB_HOST"),
			Driver:       os.Getenv("DB_DRIVER"),
			Database:     os.Getenv("DB_NAME"),
			Port:         os.Getenv("DB_PORT"),
			User:         os.Getenv("DB_USER"),
			Password:     os.Getenv("DB_PASSWORD"),
			LogLevel:     os.Getenv("LOG_LEVEL"),
			JWTSecret:    []byte(os.Getenv("JWT_SECRET")),
			AuthCookie:   os.Getenv("AUTH_COOKIE"),
			AllowedHosts: os.Getenv("ALLOWED_HOSTS"),
			RedisPort:    os.Getenv("REDIS_PORT"),
		}
	}

	return config

}
