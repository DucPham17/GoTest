FROM golang:latest

LABEL maintainer="Quique <phamminhduckl@gmail.com>"

WORKDIR /app

COpy go.mod .

copy go.sum .

run go mod download

copy . .

run go build 

run find . -name "*.go" -type f -delete

expose 3000

cmd ["./gotest"]