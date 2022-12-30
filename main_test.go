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

	as := Account{}
	createdAccData, err := Create(as, &accRequest)

	if err != nil {
		t.Error("error should not be nil")
	}

	id = createdAccData.ID
	accData, err := Get(as, id)

	if err != nil {
		t.Error("error should not be nil")
	}

	if (AccountData{}) == accData {
		t.Errorf("account data should not be nil, but was %v", accData)
	}

	if accData.ID == "" {
		t.Errorf("account id should not be nil, but was %s", accData.ID)
	}
}
