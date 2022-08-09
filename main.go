package main

import (
	"log"
	"mitramas_test/routes"

	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatalf("Error read env file with err: %s", errEnv)
	}
	// db.Connect()
	routes.Init()
}
