package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_CONN_STR string
)

func InitEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	DB_CONN_STR = os.Getenv("DB_CONN_STR")
}
