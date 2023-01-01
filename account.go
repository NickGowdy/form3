package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const resource = "organisation/accounts"

type Account struct {
	Id                   string
	Version              int64
	AccountCreateRequest AccountCreateRequest
}

// DoFetch calls the Form3 API and returns an account response
// containing account bank details.
//
// An account is fetched using the account Id and account version.
// If the account isn't fetched, an error will be returned instead
// and the account response struct will be nil.
func DoFetch(f Form3) (AccountResponse, error) {
	resp, err := f.fetch()
	return decode(err, resp)
}

// DoCreate calls the Form3 API and returns an account response
// containing account bank details.
//
// A new account is saved using an Account struct with an
// AccountCreateRequest. If creating an account is unsuccessful,
// an error will be returned with the missing values, for example:
// "validation failure list:\nvalidation failure list:\nattributes in body is required"
func DoCreate(f Form3) (AccountResponse, error) {
	resp, err := f.create()
	return decode(err, resp)
}

// DoDelete calls the Form3 API and returns a bool
// for successful deletion.
//
// An account is deleted using the account Id and account version.
// If the account isn't deleted, an error will be returned instead.
func DoDelete(f Form3) (bool, error) {
	resp, err := f.delete()

	return resp.StatusCode == http.StatusNoContent, err
}

func (a Account) fetch() (http.Response, error) {
	url := fmt.Sprintf("%s/%s/%s", os.Getenv("BASE_URL"), resource, a.Id)
	resp, err := http.Get(url)

	return *resp, err
}

func (a Account) create() (http.Response, error) {
	url := fmt.Sprintf("%s/%s/%s", os.Getenv("BASE_URL"), resource, a.Id)
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
	url := fmt.Sprintf("%s/%s/%s?version=%v", os.Getenv("BASE_URL"), resource, a.Id, a.Version)
	req, err := http.NewRequest(http.MethodDelete, url, nil)

	if err != nil {
		return http.Response{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if resp.StatusCode == http.StatusNotFound {
		return *resp, fmt.Errorf(fmt.Sprintf("record %s does not exist", a.Id))
	}

	return *resp, err
}

func decode(err error, resp http.Response) (AccountResponse, error) {
	var acc AccountResponse
	var accErr AccountError

	switch resp.StatusCode {
	case 400, 404, 409:
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
