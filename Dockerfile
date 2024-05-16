FROM golang:1.22.3-alpine AS builder
LABEL authors="PATRICK-ANJOS"

RUN apk add bash

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o api cmd/api/main.go


FROM alpine:latest

COPY --from=builder /app .

# Variáveis de ambiente para conexão com o MongoDB
# Variável de ambiente para a URL de conexão com o MongoDB
ENV MONGODB_URL=from_compose

ARG DEFAULT_PORT=from_compose
ENV API_PORT $DEFAULT_PORT

EXPOSE $API_PORT

RUN chmod +x api

CMD ["./api"]