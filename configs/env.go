package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
    if err := godotenv.Load("local.env"); err != nil {
        log.Fatal("Error load local.env file", err)
    }
}