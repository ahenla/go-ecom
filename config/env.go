package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAdress   string
	DBName     string
}

var Envs = InitConfig()

func InitConfig() Config {

	godotenv.Load()

	return Config{
		PublicHost: getENV("PUBLIC_HOST", "http://localhost"),
		Port:       getENV("PORT", "8080"),
		DBUser:     getENV("DB_USER", "root"),
		DBPassword: getENV("DB_PASSWORD", "C0ntr4s3n4-"),
		DBAdress:   fmt.Sprintf("%s:%s", getENV("DB_HOST", "127.0.0.1"), getENV("DB_PORT", "3306")),
		DBName:     getENV("DB_NAME", "ecomerce_go"),
	}
}

func getENV(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
