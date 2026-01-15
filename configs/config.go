package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port PortConfig
}

type PortConfig struct {
	name string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file, using default settiongs")
	}
	return &Config{
		Port: PortConfig{
			name: os.Getenv("PORT"),
		},
	}
}
