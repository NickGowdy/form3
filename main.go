package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type API interface {
	fetch() (http.Response, error)
	delete() (http.Response, error)
}

type AccountService struct {
	id uuid.UUID
}

func (as AccountService) fetch() (http.Response, error) {
	url := fmt.Sprintf("http://localhost:8080/v1/organisation/accounts/%s", as.id)
	fmt.Println(url)
	resp, err := http.Get(url)

	return *resp, err
}

// func (a AccountService) fetch() (AccountData, error) {
// 	url := fmt.Sprintf("http://localhost:8080/v1/organisation/accounts/%s", a.id)
// 	fmt.Println(url)
// 	resp, err := http.Get(url)
// 	var acc AccountResponse
// 	if err != nil {
// 		return *acc.AccountData, nil
// 	}

// 	if err = json.NewDecoder(resp.Body).Decode(&acc); err != nil {
// 		return *acc.AccountData, nil
// 	}

// 	return *acc.AccountData, nil
// }

func main() {
	// id, err := uuid.Parse("89faf3cd-fc6e-4e87-b930-00c182cafb05")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// as := AccountService{id: id}
	// as.fetch()
}
