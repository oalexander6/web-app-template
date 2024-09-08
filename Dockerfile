FROM node:22-alpine AS clientbuild

WORKDIR /app

COPY ./web ./web

# TODO

FROM golang:1.22 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -o ./webapp ./cmd/main.go

# Stage 2
FROM alpine:latest

WORKDIR /app/

COPY --from=1 /app/webapp ./

EXPOSE 8000

CMD ["./webapp"]