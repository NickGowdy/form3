FROM golang:alpine
ENV CGO_ENABLED=0

WORKDIR /app
COPY . .

RUN go build -o clientlibrary .

CMD [ "./clientlibrary" ]