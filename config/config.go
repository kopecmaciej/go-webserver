package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	GlobalConfig Config
)

type Config struct {
	Server   Server
	Database Database
	Cache    Cache
}

type Server struct {
	Port string
}

type Database struct {
	Url string
}

type Cache struct {
	Url      string
	User     string
	Password string
}

func getEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func LoadConfig() {
	GlobalConfig = Config{
		Server: Server{
			Port: getEnv("APP_PORT"),
		},
		Database: Database{
			Url: getEnv("DATABASE_URL"),
		},
		Cache: Cache{
			Url:      getEnv("CACHE_URL"),
			User:     getEnv("CACHE_USER"),
			Password: getEnv("CACHE_PASSWORD"),
		},
	}
}
