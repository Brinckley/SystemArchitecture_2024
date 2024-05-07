FROM golang:latest

RUN mkdir /u_app
WORKDIR /u_app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . /u_app
WORKDIR /u_app/cmd/app
RUN go build -o main

CMD ["./main"]