FROM golang:1.19-bullseye AS builder

WORKDIR /go/src/github.com/mercari/mercari-microservices-example

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
RUN go build -o /usr/bin/catalog ./services/catalog

CMD ["/usr/bin/catalog"]
