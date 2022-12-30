package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const endpoint = "organisation/accounts"

type API interface {
	create(*AccountCreateRequest) (http.Response, error)
	fetch(id string) (http.Response, error)
	delete(id string, version int64) (http.Response, error)
}

type Account struct {
	AccountResponse AccountResponse
}

// create implements API
func (Account) create(ar *AccountCreateRequest) (http.Response, error) {
	url := fmt.Sprintf("http://localhost:8080/v1/%s", endpoint)
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(&ar)
	if err != nil {
		return http.Response{}, err
	}

	resp, err := http.Post(url, "application/json", b)
	if err != nil {
		return http.Response{}, err
	}

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(body)

	return *resp, err
}

// delete implements API
func (Account) delete(id string, version int64) (http.Response, error) {
	url := fmt.Sprintf("http://localhost:8080/v1/%s/%s?version=%v", endpoint, id, version)
	req, err := http.NewRequest(http.MethodDelete, url, nil)

	if err != nil {
		return http.Response{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	return *resp, err
}

func (as Account) fetch(id string) (http.Response, error) {
	url := fmt.Sprintf("http://localhost:8080/v1/%s/%s", endpoint, id)
	resp, err := http.Get(url)

	return *resp, err
}

func Fetch(api API, id string) (AccountResponse, error) {
	resp, err := api.fetch(id)
	return decode(err, resp)
}

func Create(api API, ar *AccountCreateRequest) (AccountResponse, error) {
	resp, err := api.create(ar)

	return decode(err, resp)
}

func decode(err error, resp http.Response) (AccountResponse, error) {
	var acc AccountResponse
	if err != nil {
		return acc, nil
	}

	if err = json.NewDecoder(resp.Body).Decode(&acc); err != nil {
		return acc, nil
	}

	defer resp.Body.Close()

	return acc, nil
}

func Delete(api API, id string, version int64) (bool, error) {
	resp, err := api.delete(id, version)
	if err != nil {
		log.Fatal(err)
	}

	return resp.StatusCode == http.StatusNoContent, err
}

func main() {
}
