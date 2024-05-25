package main

import (
	"fmt"
	"gophkeeper/internal/client/app"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	// Version Версия сборки
	Version string
	// BuildDate Дата сборки
	BuildDate string
)

func main() {
	fmt.Println("Start client...")
	fmt.Println("Version:", Version)
	fmt.Println("Build Date:", BuildDate)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("No .env file found")
	}

	config := map[string]string{
		"SERVER_ADDRESS": os.Getenv("SERVER_ADDRESS"),
		"DB_FILE":        os.Getenv("DB_FILE"),
	}

	app.New(config).Start()
}
