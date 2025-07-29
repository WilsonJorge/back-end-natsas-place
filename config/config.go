package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl        string
	JWTSecret          string
	JWTExpirationHours int
	EmailSender        string
	EmailPassword      string
	SMTPHost           string
	SMTPPort           string
	MailgunDomain      string
	MailgunAPIKey      string
	StripeSecretKey    string
	StripePublicKey    string
	FrontendURL        string
	BackendURL         string
	Environment        string
	Port               string
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
		DatabaseUrl:        os.Getenv("DATABASE_URL"),
		JWTSecret:          os.Getenv("JWT_SECRET"),
		JWTExpirationHours: 24, // Default 24 hours
		// EmailSender:        os.Getenv("EMAIL_SENDER"),
		// EmailPassword:      os.Getenv("EMAIL_PASSWORD"),
		// SMTPHost:           os.Getenv("SMTP_HOST"),
		// SMTPPort:           os.Getenv("SMTP_PORT"),
		// MailgunDomain:      os.Getenv("MAILGUN_DOMAIN"),
		// MailgunAPIKey:      os.Getenv("MAILGUN_API_KEY"),
		// StripeSecretKey:    os.Getenv("STRIPE_SECRET_KEY"),
		// StripePublicKey:    os.Getenv("STRIPE_PUBLIC_KEY"),
		FrontendURL: os.Getenv("FRONTEND_URL"),
		BackendURL:  os.Getenv("BACKEND_URL"),
		Environment: os.Getenv("ENVIRONMENT"),
		Port:        os.Getenv("PORT"),
	}

	if config.DatabaseUrl == "" {
		log.Fatal("DATABASE_URL não encontrada nas variáveis de ambiente")
	}

	if config.JWTSecret == "" {
		log.Fatal("JWT_SECRET não encontrada nas variáveis de ambiente")
	}

	if config.Port == "" {
		config.Port = "3001"
	}

	if config.Environment == "" {
		config.Environment = "development"
	}

	if config.FrontendURL == "" {
		config.FrontendURL = "http://localhost:3000"
	}

	if config.BackendURL == "" {
		config.BackendURL = "http://localhost:3001"
	}

	return config
}
