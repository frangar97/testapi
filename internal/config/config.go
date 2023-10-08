package config

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	Port         string
	Host         string
	User         string
	Password     string
	Database     string
	DatabasePort string
}

func InitConfig() (config, error) {
	err := godotenv.Load()
	cfg := config{}

	if err != nil {
		return cfg, err
	}

	cfg.Port = os.Getenv("PORT")
	cfg.Database = os.Getenv("DATABASE")
	cfg.Host = os.Getenv("HOST")
	cfg.User = os.Getenv("USER")
	cfg.Password = os.Getenv("PASSWORD")
	cfg.DatabasePort = os.Getenv("DATABASE_PORT")

	return cfg, nil
}
