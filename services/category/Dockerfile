FROM golang:1.17.1-bullseye AS builder

WORKDIR /go/src/github.com/mercari/go-conference-2021-spring-office-hour

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
RUN go build -o /usr/bin/category ./services/category

CMD ["/usr/bin/category"]
