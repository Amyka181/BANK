FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
COPY /config/config.yaml /app/config.yaml

RUN  go mod download

COPY . .

WORKDIR /app/cmd/apiserver

RUN  go build -o Bankirka .

EXPOSE 8080
EXPOSE 50051

CMD ["./Bankirka"]