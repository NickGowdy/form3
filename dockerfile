FROM golang:1.12-alpine

RUN apk add alpine-sdk
ENV CGO_ENABLED=0

# Set the Current Working Directory inside the container
WORKDIR /app/form3-clientlibrary

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./clientlibrary .

# Run the binary program produced by `go install`
CMD ["./clientlibrary"]

