package main

import (
	"testing"

	"github.com/google/uuid"
)

func TestCreateFetchDelete(t *testing.T) {
	id := uuid.NewString()
	orgId := uuid.NewString()
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

	accRequest := AccountCreateRequest{
		AccountData: &AccountData{
			ID:             id,
			OrganisationID: orgId,
			Type:           "accounts",
			Attributes:     &accAttributes,
		}}

	as := Account{AccountCreateRequest: accRequest}
	createdAccResp, err := DoCreate(as)

	if err != nil {
		t.Error("error should not be nil")
	}

	as.Id = createdAccResp.AccountData.ID
	as.Version = *createdAccResp.AccountData.Version

	accData, err := DoFetch(as)

	if err != nil {
		t.Error("error should not be nil")
	}

	if (AccountResponse{}) == accData {
		t.Errorf("account data should not be nil, but was %v", accData)
	}

	if accData.AccountData.ID == "" {
		t.Errorf("account id should not be nil, but was %s", accData.AccountData.ID)
	}

	isDeleted, err := DoDelete(as)

	if err != nil {
		t.Error("error should not be nil")
	}

	if isDeleted != true {
		t.Errorf("is deleted should be true, but was %v", isDeleted)
	}
}
