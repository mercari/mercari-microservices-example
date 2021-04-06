FROM golang:1.16.3-buster AS builder

ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/mercari/go-conference-2021-spring-office-hour

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
RUN go build -o /usr/bin/balance ./services/balance

## Runtime

FROM gcr.io/distroless/base:1e4a8bb3ad03f71b572cbcb3bbc25f3fd8d0ff14
COPY --from=builder /usr/bin/balance /usr/bin/balance

CMD ["/usr/bin/balance"]
