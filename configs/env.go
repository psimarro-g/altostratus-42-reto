package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMongoURI(ENV string) string {
	if ENV == "local" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	return os.Getenv("MONGOURI")
}
