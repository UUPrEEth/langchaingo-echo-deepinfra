package coreSetup

import (
	"fmt"
	"log"
	"os"
)

func GetEnv(env string) (string, error) {
	envVariable := os.Getenv(env)
	if env == "" {
		log.Fatal(env + " environment variable not set")
		return "", fmt.Errorf("%s environment variable not set", env)
	}
	return envVariable, nil
}
