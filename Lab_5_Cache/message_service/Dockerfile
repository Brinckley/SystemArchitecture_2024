FROM golang:latest

RUN mkdir /m_app
WORKDIR /m_app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . /m_app
WORKDIR /m_app/cmd/app
RUN go build -o main

EXPOSE 8088

CMD ["./main"]