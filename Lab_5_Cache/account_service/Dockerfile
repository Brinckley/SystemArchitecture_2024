FROM golang:latest

RUN mkdir /a_app
WORKDIR /a_app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . /a_app
WORKDIR /a_app/cmd/app
RUN go build -o main

EXPOSE 8087

CMD ["./main"]