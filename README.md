# Form3 API Client Library
<p>This is a client library used to access the Form3 Account API. It provides a common interface for fetching, creating and deleting account records.</p>

## How to run this client library
---

The recommended way of running this client library on any machine locally is through Docker.

This can be done with the following steps:

- Download Docker [here](https://www.docker.com/), select your OS and install it to your local machine.
- Verify Docker is running via this command: `docker -v` (You should see your Docker version).
- Type the command `Docker compose up -d` which will load all the service integrations that you need running locally.
- Then type `Docker compose run clientlibrary` which will run the Form3 Client Library with it's unit tests to confirm the integrations are configured correctly.

