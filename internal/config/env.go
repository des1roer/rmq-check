package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env") // Загружаем основной файл в любом случае

	// Проверяем существует ли .env.local перед загрузкой
	if _, err := os.Stat(".env.local"); err == nil {
		if err := godotenv.Overload(".env.local"); err != nil {
			log.Fatalf("Error loading .env.local file: %v", err)
		}
	}

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func init() {
	if fileExists(".env") {
		LoadEnv()
	}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		// Если произошла другая ошибка (например, нет прав доступа), можно обработать её
		log.Printf("Error checking file: %v", err)
		return false
	}
	return true
}
