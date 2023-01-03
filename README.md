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