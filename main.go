package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var host = os.Getenv("HOST")
var port = os.Getenv("PORT")

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error Loading Env Variables")
	}
}
