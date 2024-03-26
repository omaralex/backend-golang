package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func ReadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	appPort, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		panic(err.Error())
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err.Error())
	}

	return &Config{
		Server: ServerConfig{
			Version: os.Getenv("APP_VERSION"),
			Port:    appPort,
		},
		Security: SecurityConfig{
			JWTSecret: os.Getenv("JWT_SECRET"),
		},
		PostgreSql: PostgreSqlConfig{
			Server:   os.Getenv("DB_SERVER"),
			Port:     dbPort,
			Database: os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
		},
	}
}
