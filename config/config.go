package config

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	DBHost string `validate:"required,hostname|ip"`
	DBPort string `validate:"required,numeric"`
	DBUser string `validate:"required"`
	DBPass string `validate:"required"`
	DBName string `validate:"required"`
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	var validate = validator.New()

	cfg := &Config{
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBName: os.Getenv("DB_NAME"),
	}
	if err := validate.Struct(cfg); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldErr := range validationErrors {
			log.Printf("Validation error of field %s", fieldErr.Field())
		}
		log.Fatal("There's an error in configuration")
	}

	return cfg
}
