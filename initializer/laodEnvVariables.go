package initializer

import (
	"log"
	"github.com/joho/godotenv"
)

func LoadEnvVaraibles()  {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
}