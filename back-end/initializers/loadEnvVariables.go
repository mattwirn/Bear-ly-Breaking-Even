package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables(path string) {
	err := godotenv.Load(path)

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
