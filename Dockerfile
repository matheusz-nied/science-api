# syntax=docker/dockerfile:1
FROM golang:1.22.5-alpine

# Instalar o git e bash
RUN apk add --no-cache git bash

WORKDIR /app

# Instalar o air
RUN go install github.com/air-verse/air@latest

# Copia o go.mod e go.sum para o contêiner
COPY go.mod go.sum ./
RUN go mod download

# Copia todo o código fonte para o contêiner
COPY . ./

# Expõe a porta 8079
EXPOSE 8079

# Define o comando de inicialização do contêiner
CMD ["air", "-c", ".air.toml"]
