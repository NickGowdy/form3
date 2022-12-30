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
	fetch() (http.Response, error)
	delete() (http.Response, error)
}

type Account struct {
	Id      string
	Version int64
	Error   error
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
func (a Account) delete() (http.Response, error) {
	url := fmt.Sprintf("http://localhost:8080/v1/%s/%s?version=%v", endpoint, a.Id, a.Version)
	req, err := http.NewRequest(http.MethodDelete, url, nil)

	if err != nil {
		return http.Response{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	return *resp, err
}

func (a Account) fetch() (http.Response, error) {
	url := fmt.Sprintf("http://localhost:8080/v1/%s/%s", endpoint, a.Id)
	resp, err := http.Get(url)

	return *resp, err
}

func DoFetch(api API) (AccountResponse, error) {
	resp, err := api.fetch()
	return decode(err, resp)
}

func DoCreate(api API, ar *AccountCreateRequest) (AccountResponse, error) {
	resp, err := api.create(ar)

	return decode(err, resp)
}

func DoDelete(api API) (bool, error) {
	resp, err := api.delete()
	if err != nil {
		log.Fatal(err)
	}

	return resp.StatusCode == http.StatusNoContent, err
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

func main() {
}
