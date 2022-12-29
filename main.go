package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

const endpoint = "organisation/accounts"

type API interface {
	fetch(id uuid.UUID) (http.Response, error)
	delete(id uuid.UUID) (http.Response, error)
}

type Account struct {
}

// delete implements API
func (Account) delete(id uuid.UUID) (http.Response, error) {
	panic("unimplemented")
}

func (as Account) fetch(id uuid.UUID) (http.Response, error) {
	url := fmt.Sprintf("http://localhost:8080/v1/%s/%s", endpoint, id)
	fmt.Println(url)
	resp, err := http.Get(url)

	return *resp, err
}

func Get(api API, id uuid.UUID) (AccountData, error) {
	resp, err := api.fetch(id)
	var acc AccountResponse
	if err != nil {
		return *acc.AccountData, nil
	}

	if err = json.NewDecoder(resp.Body).Decode(&acc); err != nil {
		return *acc.AccountData, nil
	}

	return *acc.AccountData, nil
}

func main() {
	id, err := uuid.Parse("89faf3cd-fc6e-4e87-b930-00c182cafb05")
	if err != nil {
		log.Fatal(err)
	}

	as := Account{}
	Get(as, id)
}
