FROM golang:latest

RUN mkdir /app
WORKDIR /app

CMD ["ls"]
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . /app
WORKDIR /app/cmd/app
RUN go build -o main

CMD ["./main"]