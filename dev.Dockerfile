FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go install github.com/githubnemo/CompileDaemon@latest

COPY cmd ./cmd
COPY internal ./internal
COPY pkg ./pkg
COPY config ./config
COPY test ./test

RUN go build -o /bin/db ./cmd/db/
