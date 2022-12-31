package main

import "os"

func main() {

	if os.Getenv("BASE_URL") == "" {
		panic("BASE_URL is not set")
	}
}
