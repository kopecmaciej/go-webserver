# syntax=docker/dockerfile:1

ARG GO_V=1.18.3

FROM golang:${GO_V}-alpine3.16 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./web-server ./cmd/main.go

FROM alpine:3.16 

COPY --from=build /app/web-server ./
COPY --from=build /app/.env /.env

EXPOSE 4000

ENTRYPOINT ["./web-server"]

