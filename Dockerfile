FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN  go mod download

COPY . .

WORKDIR /app/cmd/apiserver

RUN  go build -o Bankirka .

EXPOSE 8080

CMD ["./Bankirka"]