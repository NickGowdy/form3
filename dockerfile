# syntax=docker/dockerfile:1
FROM golang:1.19.4-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /form3

EXPOSE 8081

CMD ["/form3"]