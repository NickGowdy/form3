package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/nickgowdy/form3/account"
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
	as := account.NewPing(30)
	resp, err := account.DoPing(as)

	if err != nil || !resp {
		panic(err)
	}
}
