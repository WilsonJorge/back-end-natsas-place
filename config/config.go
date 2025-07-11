package config

import (
	"log"
	"os"
	"sync"

	"github.com/lpernett/godotenv"
)

type Config struct {
	DatabaseUrl string
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		instance = Init()
	})
	return instance
}

func Init() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &Config{
		DatabaseUrl: os.Getenv("DATABASE_URL"),
	}

	if config.DatabaseUrl == "" {
		log.Fatal("DATABASE_URL não encontrada nas variáveis de ambiente")
	}

	return config
}
