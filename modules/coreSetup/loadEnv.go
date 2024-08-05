package coreSetup

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariable(env string) {

	// load .env file
	err := godotenv.Load("./env/" + env + ".env")
	if err != nil {
		log.Fatal(err.Error())
		log.Fatalf("Error loading .env file")
	}
}
