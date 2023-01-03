package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// some checks to make sure library is initialised correctly
	checkEnv()
	pingServer()
}

func checkEnv() {
	godotenv.Load()

	if os.Getenv("BASE_URL") == "" {
		panic("BASE_URL is not set")
	}
}

func pingServer() {
	as := Account{}
	resp, err := as.ping()

	if err != nil || !resp {
		panic(err)
	}
}
