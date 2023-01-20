package account

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/joho/godotenv"
)

func TestCreateFetchDelete(t *testing.T) {
	os.Setenv("BASE_URL", "http://localhost:8080/v1")

	country := "GB"
	accClassification := "Personal"
	accAttributes := AccountAttributes{
		AccountClassification: &accClassification,
		AccountNumber:         "10000004",
		BankID:                "400302",
		BankIDCode:            "GBDSC",
		BaseCurrency:          "GBP",
		Bic:                   "NWBKGB42",
		Country:               &country,
		Iban:                  "GB28NWBK40030212764204",
		JointAccount:          new(bool),
		Name:                  []string{"Nick", "Gowdy"},
	}

	account := NewCreateAccount(0, "accounts", &accAttributes, 30)
	createdAccResp, err := DoCreate(account)

	if err != nil {
		t.Errorf("error should be nil, but is: %s", err)
	}

	account = NewFetchAccount(createdAccResp.AccountData.ID, *createdAccResp.AccountData.Version, 30)
	accResp, err := DoFetch(account)

	if err != nil {
		t.Errorf("error should be nil, but is: %s", err)
	}

	if accResp == nil {
		t.Errorf("account response should not nil")
	}

	if !reflect.DeepEqual(createdAccResp.AccountData.Attributes.AccountNumber, "10000004") {
		t.Errorf("got %v want %v", createdAccResp.AccountData.Attributes.AccountNumber, "10000004")
	}

	if !reflect.DeepEqual(createdAccResp.AccountData.Attributes.BankID, "400302") {
		t.Errorf("got %v want %v", createdAccResp.AccountData.Attributes.BankID, "400302")
	}

	if !reflect.DeepEqual(createdAccResp.AccountData.Attributes.BankIDCode, "GBDSC") {
		t.Errorf("got %v want %v", createdAccResp.AccountData.Attributes.BankIDCode, "GBDSC")
	}

	account = NewDeleteAccount(createdAccResp.AccountData.ID, *createdAccResp.AccountData.Version, 30)
	isDeleted, err := DoDelete(account)

	if err != nil {
		t.Errorf("error should be nil, but is: %s", err)
	}

	if isDeleted != true {
		t.Errorf("is deleted should be true, but was %v", isDeleted)
	}
}

func TestFetchAccountDontExist(t *testing.T) {
	godotenv.Load()

	id := generateId()
	account := NewFetchAccount(id, 0, 30)
	_, err := DoFetch(account)

	expected := fmt.Sprintf("record %s does not exist", id)
	if fmt.Sprint(err) != expected {
		t.Errorf("error message should be: %s", expected)
	}
}

func TestDeleteAccountDontExist(t *testing.T) {
	godotenv.Load()

	id := generateId()
	account := NewDeleteAccount(id, 0, 30)
	_, err := DoDelete(account)

	expected := fmt.Sprintf("record %s does not exist", id)
	if fmt.Sprint(err) != expected {
		t.Errorf("error message should be: %s", expected)
	}
}

func TestCreateInvalidAccountDataFields(t *testing.T) {
	godotenv.Load()

	id := generateId()
	orgId := generateId()
	accAttributes := AccountAttributes{}

	accRequest := accountCreateRequest{
		AccountData: &accountData{}}

	account := Account{AccountCreateRequest: accRequest}
	_, err := DoCreate(account)

	expected := "validation failure list:\nvalidation failure list:\nattributes in body is required\nid in body is required\norganisation_id in body is required\ntype in body is required"

	if fmt.Sprint(err) != expected {
		t.Errorf("error message should be: %s", expected)
	}

	account.AccountCreateRequest.AccountData.ID = id
	_, err = DoCreate(account)
	expected = "validation failure list:\nvalidation failure list:\nattributes in body is required\norganisation_id in body is required\ntype in body is required"

	if fmt.Sprint(err) != expected {
		t.Errorf("error message should be: %s", expected)
	}

	account.AccountCreateRequest.AccountData.OrganisationID = orgId
	_, err = DoCreate(account)
	expected = "validation failure list:\nvalidation failure list:\nattributes in body is required\ntype in body is required"

	if fmt.Sprint(err) != expected {
		t.Errorf("error message should be: %s", expected)
	}

	account.AccountCreateRequest.AccountData.Type = "accounts"
	_, err = DoCreate(account)
	expected = "validation failure list:\nvalidation failure list:\nattributes in body is required"

	if fmt.Sprint(err) != expected {
		t.Errorf("error message should be: %s", expected)
	}

	account.AccountCreateRequest.AccountData.Attributes = &accAttributes
	_, err = DoCreate(account)
	expected = "validation failure list:\nvalidation failure list:\nvalidation failure list:\ncountry in body is required\nname in body is required"

	if fmt.Sprint(err) != expected {
		t.Errorf("error message should be: %s", expected)
	}
}

func TestCreateInvalidAccountAttributeFields(t *testing.T) {
	godotenv.Load()

	id := generateId()
	orgId := generateId()
	country := "GB"
	name := []string{"Nick", "Gowdy"}

	accRequest := accountCreateRequest{
		AccountData: &accountData{
			ID:             id,
			OrganisationID: orgId,
			Type:           "accounts",
			Attributes:     &AccountAttributes{},
		}}

	account := Account{AccountCreateRequest: accRequest}
	_, err := DoCreate(account)

	expected := "validation failure list:\nvalidation failure list:\nvalidation failure list:\ncountry in body is required\nname in body is required"

	if fmt.Sprint(err) != expected {
		t.Errorf("error message should be: %s", expected)
	}

	accRequest.AccountData.Attributes.Country = &country
	_, err = DoCreate(account)
	expected = "validation failure list:\nvalidation failure list:\nvalidation failure list:\nname in body is required"

	if fmt.Sprint(err) != expected {
		t.Errorf("error message should be: %s", expected)
	}

	accRequest.AccountData.Attributes.Name = name
	accResp, err := DoCreate(account)
	if err != nil {
		t.Error("error should be nil")
	}

	if accResp == nil {
		t.Errorf("account response should not nil")
	}
}

func TestCreateDuplicateAccount(t *testing.T) {
	godotenv.Load()

	id := generateId()
	orgId := generateId()
	country := "GB"
	accClassification := "Personal"
	accAttributes := AccountAttributes{
		AccountClassification:   &accClassification,
		AccountNumber:           "10000004",
		BankID:                  "400302",
		BankIDCode:              "GBDSC",
		BaseCurrency:            "GBP",
		Bic:                     "NWBKGB42",
		Country:                 &country,
		Iban:                    "GB28NWBK40030212764204",
		JointAccount:            new(bool),
		Name:                    []string{"Nick", "Gowdy"},
		SecondaryIdentification: id,
	}

	accRequest := accountCreateRequest{
		AccountData: &accountData{
			ID:             id,
			OrganisationID: orgId,
			Type:           "accounts",
			Attributes:     &accAttributes,
		}}

	account := Account{AccountCreateRequest: accRequest}
	_, err := DoCreate(account)

	if err != nil {
		t.Errorf("error should be nil, but is: %s", err)
	}

	_, err = DoCreate(account)

	expected := "Account cannot be created as it violates a duplicate constraint"
	if fmt.Sprint(err) != expected {
		t.Errorf("error message should be: %s", expected)
	}
}
