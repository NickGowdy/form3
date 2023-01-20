package main

import (
	"encoding/json"
	"log"

	"github.com/joho/godotenv"
	"github.com/nickgowdy/form3/account"
)

func main() {

	godotenv.Load()

	country := "GB"
	accClassification := "Personal"

	accAttributes := account.AccountAttributes{
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

	acc := account.NewCreateAccount(1, "accounts", &accAttributes, 30)

	err := account.DoPing(acc)

	if err != nil {
		log.Fatal("error connecting to API:", err)
	}

	accResponse, err := account.DoCreate(acc)

	if err != nil {
		log.Fatal("error creating account", err)
	}

	j, _ := json.Marshal(accResponse)
	log.Printf("Account %+v", string(j))
}
