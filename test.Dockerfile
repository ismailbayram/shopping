FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go install honnef.co/go/tools/cmd/staticcheck@latest


COPY cmd ./cmd
COPY internal ./internal
COPY pkg ./pkg
COPY config ./config
COPY test ./test

RUN chmod +x ./test/test_and_coverage.sh

ENTRYPOINT [ "./test/test_and_coverage.sh" ]