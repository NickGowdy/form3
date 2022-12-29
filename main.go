package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080/v1/organisation/accounts/89faf3cd-fc6e-4e87-b930-00c182cafb05")
	if err != nil {
		log.Fatal(err)
	}

	var acc AccountResponse
	if err = json.NewDecoder(resp.Body).Decode(&acc); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Account is: %v", *acc.AccountData.Attributes.Country)
}
