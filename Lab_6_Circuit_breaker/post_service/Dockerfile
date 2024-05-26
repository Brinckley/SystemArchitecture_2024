FROM golang:latest

RUN mkdir /p_app
WORKDIR /p_app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . /p_app
WORKDIR /p_app/cmd/app
RUN go build -o main

EXPOSE 8089

CMD ["./main"]