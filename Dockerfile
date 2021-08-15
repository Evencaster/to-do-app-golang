FROM golang:1.16.7-alpine3.13

WORKDIR /src/to-do-app-golang

COPY . .

RUN go get -d ./...

RUN go run main.go
