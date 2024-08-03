package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAdress               string
	DBName                 string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

var Envs = InitConfig()

func InitConfig() Config {

	godotenv.Load()

	return Config{
		PublicHost:             getENV("PUBLIC_HOST", "http://localhost"),
		Port:                   getENV("PORT", "8080"),
		DBUser:                 getENV("DB_USER", "root"),
		DBPassword:             getENV("DB_PASSWORD", "C0ntr4s3n4-"),
		DBAdress:               fmt.Sprintf("%s:%s", getENV("DB_HOST", "127.0.0.1"), getENV("DB_PORT", "3306")),
		DBName:                 getENV("DB_NAME", "ecomerce_go"),
		JWTSecret:              getENV("JWT_SECRET", "not-so-secret"),
		JWTExpirationInSeconds: getENVAsInt("JWT_EXP", 3600*24*7),
	}
}

func getENV(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getENVAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}
	return fallback
}
