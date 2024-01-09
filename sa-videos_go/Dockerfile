FROM golang:latest

WORKDIR /usr/src/app

COPY . .

RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon
RUN go mod tidy
