FROM golang:1.20-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY pkg ./pkg
COPY config ./config
COPY test ./test

RUN go build -o /app/db ./cmd/db/
RUN go build -o /app/shopping-server ./cmd/server/

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/config /app/config
COPY --from=build /app/shopping-server /app/shopping-server
COPY --from=build /app/db /app/db

CMD ["./shopping-server"]