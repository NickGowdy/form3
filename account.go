package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const endpoint = "organisation/accounts"

type Account struct {
	Id                   string
	Version              int64
	AccountCreateRequest AccountCreateRequest
}

func DoFetch(f Form3) (AccountResponse, error) {
	resp, err := f.fetch()
	return decode(err, resp)
}

func DoCreate(f Form3) (AccountResponse, error) {
	resp, err := f.create()
	return decode(err, resp)
}

func DoDelete(f Form3) (bool, error) {
	resp, err := f.delete()
	if err != nil {
		log.Fatal(err)
	}
	return resp.StatusCode == http.StatusNoContent, err
}

func (a Account) fetch() (http.Response, error) {
	url := fmt.Sprintf("http://localhost:8080/v1/%s/%s", endpoint, a.Id)
	resp, err := http.Get(url)

	return *resp, err
}

func (a Account) create() (http.Response, error) {
	url := fmt.Sprintf("http://localhost:8080/v1/%s", endpoint)
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(&a.AccountCreateRequest)
	if err != nil {
		return http.Response{}, err
	}

	resp, err := http.Post(url, "application/json", b)
	if err != nil {
		return http.Response{}, err
	}

	return *resp, err
}

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

func decode(err error, resp http.Response) (AccountResponse, error) {
	var acc AccountResponse
	var accErr AccountError

	switch resp.StatusCode {
	case 400, 409:
		if err = json.NewDecoder(resp.Body).Decode(&accErr); err != nil {
			return acc, err
		}

		return acc, fmt.Errorf(accErr.ErrorMessage)
	}

	if err != nil {
		return acc, nil
	}

	if err = json.NewDecoder(resp.Body).Decode(&acc); err != nil {
		return acc, nil
	}

	defer resp.Body.Close()

	return acc, nil
}
