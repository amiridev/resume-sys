package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func InitEnvs() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetRunAddress() string {
	return "0.0.0.0:" + GetEnv("APP_PORT")
}

func GetDatabaseDNS() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		GetEnv("DB_USERNAME"),
		GetEnv("DB_PASSWORD"),
		GetEnv("DB_HOST"),
		GetEnv("DB_PORT"),
		GetEnv("DB_NAME"),
	)
}

func GetTokensExpires() (accessExpiresHours, refreshExpiresDays int) {
	accessExpires, _ := strconv.Atoi(GetEnv("TOKEN_ACCESS_EXPIRES_HOURS"))
	accessExpiresHours = accessExpires

	refreshExpires, _ := strconv.Atoi(GetEnv("TOKEN_REFRESH_EXPIRES_DAYS"))
	refreshExpiresDays = refreshExpires

	return
}
