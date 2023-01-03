# Form3 API Client Library
<p>This is a client library used to access the Form3 Account API. It provides a common interface for fetching, creating and deleting account records.</p>

## How to run this client library
---

<p>The recommended way of running this client library locally is through Docker.

This can be done with the following steps:

- Download Docker [here](https://www.docker.com/), select your OS and install it to your local machine.
- Verify Docker is running via this command: `docker -v` (You should see your Docker version).
- Type the command `docker-compose up -d` which will load all the service integrations that you need running locally.
- Then type `docker-compose run clientlibrary` which will run the Form3 Client Library with it's unit tests to confirm the integrations are configured correctly.
</p>

<p>You can also run this application locally using the Golang CLI but you will still need Docker to run the services declared in `docker-compose.
This can be done with these steps:

- Download Docker as described above.
- Go to the Golang official site [here](https://go.dev/) to download the SDK/CLI.
- Verify Go has been installed successfully with this command: `go -v` to check the Golang version.
- Create a new `.env` file in the root of the project and add this line to it: `BASE_URL=http://localhost:8080/v1`
- To run all unit tests type `go test -v ./...`
</p>

## How to run this client library
---
<p>
To create a new account using the Form3 Client Library, you will need to use `AccountCreateRequest` struct like this:

```
 AccountCreateRequest{
		AccountData: &AccountData{
			ID:             id,
			OrganisationID: orgId,
			Type:           "accounts",
			Attributes:     &accAttributes,
		}}
```

The `ID` and `OrganisationID` are string in the UUID format, for example: `ae73ad26-80a7-4976-9127-8fed11eb8a8a `

The `AccountCreateRequest` has a field `Attributes` which is a struct for the bank details. See example below:

```
AccountAttributes{
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
```

Finally the `Account` struct is used to make the account create request:
```
account := Account{AccountCreateRequest: accRequest}
	createdAccResp, err := DoCreate(account)
```

Here is the full example:
```
id := <uuid>
	orgId := <uuid>
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

	account := Account{AccountCreateRequest: accRequest}
	createdAccResp, err := DoCreate(account)
```
If the account is saved successfully, an `AccountResponse` struct will be returned with the account record.

To Fetch or Delete a record, all that is need are the following:

- Account Id
- Account Version

```
account := Account{Id: <uuid>, Version: <number_as_string>}
```

Fetch will return an `AccountResponse` and Delete will return a truthy of `true` is account was deleted successfully.

</p>