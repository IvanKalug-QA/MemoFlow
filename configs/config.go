package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port PortConfig
	Auth AuthConfig
	Db   DbConfig
}

type PortConfig struct {
	Name string
}

type AuthConfig struct {
	Secret string
}

type DbConfig struct {
	Dsn string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file, using default settiongs")
	}
	return &Config{
		Port: PortConfig{
			Name: os.Getenv("PORT"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("SECRET"),
		},
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
	}
}
