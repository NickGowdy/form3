package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

const resource = "organisation/accounts"

func GenerateId() string {
	uuid, _ := uuid.NewRandom()
	id := (uuid).String()
	return id
}

// NewCreateAccount creates an Account struct for creating a new account record
func NewCreateAccount(version int64, accType string, accAttributes *AccountAttributes, timeout int32) *Account {
	return &Account{
		Client: http.Client{Timeout: time.Duration(time.Second * time.Duration(timeout))},
		AccountCreateRequest: AccountCreateRequest{AccountData: &AccountData{
			ID:             GenerateId(),
			OrganisationID: GenerateId(),
			Version:        &version,
			Attributes:     accAttributes,
			Type:           accType,
		}},
	}
}

// NewFetchAccount creates an Account struct for fetching an account record
func NewFetchAccount(id string, version int64, timeout int32) *Account {
	return &Account{
		Client:  http.Client{Timeout: time.Duration(time.Second * time.Duration(timeout))},
		Id:      id,
		Version: version,
	}
}

// NewDeleteAccount creates an Account struct for deleting an account record
func NewDeleteAccount(id string, version int64, timeout int32) *Account {
	return &Account{
		Client:  http.Client{Timeout: time.Duration(time.Second * time.Duration(timeout))},
		Id:      id,
		Version: version,
	}
}

// NewDeleteAccount creates an Account struct for deleting an account record
func NewPing(timeout int32) *Account {
	return &Account{
		Client: http.Client{Timeout: time.Duration(time.Second * time.Duration(timeout))},
	}
}

// DoFetch calls the Form3 API and returns an account response
// containing account bank details.
//
// An account is fetched using the account Id and account version.
// If the account isn't fetched, an error will be returned instead
// and the account response struct will be nil.
func DoFetch(f Form3) (*AccountResponse, error) {
	resp, err := f.fetch()

	if err != nil {
		return nil, err
	}

	ar, err := decode(err, &resp)
	return &ar, err
}

// DoCreate calls the Form3 API and returns an account response
// containing account bank details.
//
// A new account is saved using an Account struct with an
// AccountCreateRequest. If creating an account is unsuccessful,
// an error will be returned with the missing values, for example:
// "validation failure list:\nvalidation failure list:\nattributes in body is required"
func DoCreate(f Form3) (*AccountResponse, error) {
	resp, err := f.create()

	if err != nil {
		return nil, err
	}

	ar, err := decode(err, &resp)
	return &ar, err
}

// DoDelete calls the Form3 API and returns a bool
// for successful deletion.
//
// An account is deleted using the account Id and account version.
// If the account isn't deleted, an error will be returned instead.
func DoDelete(f Form3) (bool, error) {
	resp, err := f.delete()

	if err != nil {
		return false, err
	}

	return resp.StatusCode == http.StatusNoContent, err
}

// DoPing checks that a connection is available from the API
func DoPing(f Form3) (bool, error) {
	return f.ping()
}

func (a Account) ping() (bool, error) {
	url := fmt.Sprintf("%s/%s/", os.Getenv("BASE_URL"), resource)
	resp, err := a.Client.Get(url)

	if err != nil {
		log.Println(err)
	}

	return resp.StatusCode == http.StatusOK, err
}

func (a Account) fetch() (http.Response, error) {
	url := fmt.Sprintf("%s/%s/%s", os.Getenv("BASE_URL"), resource, a.Id)
	resp, err := a.Client.Get(url)

	if err != nil {
		log.Println(err)
	}

	return *resp, err
}

func (a Account) create() (http.Response, error) {
	url := fmt.Sprintf("%s/%s/%s", os.Getenv("BASE_URL"), resource, a.Id)
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(&a.AccountCreateRequest)

	if err != nil {
		return http.Response{}, err
	}

	resp, err := a.Client.Post(url, "application/json", b)

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

	resp, err := a.Client.Do(req)

	if resp.StatusCode == http.StatusNotFound {
		return *resp, fmt.Errorf(fmt.Sprintf("record %s does not exist", a.Id))
	}

	defer resp.Body.Close()

	return *resp, err
}

func decode(err error, resp *http.Response) (AccountResponse, error) {
	defer resp.Body.Close()
	var acc AccountResponse
	var accErr AccountError

	switch statusCode := resp.StatusCode; {
	case statusCode > 399 && statusCode < 500:
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

	return acc, nil
}
