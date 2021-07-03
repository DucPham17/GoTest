set -xe

go get github.com/gin-gonic/gin
go get github.com/go-playground/validator/v10
go get github.com/tpkeeper/gin-dump

go build -o bin/application ./src/main.go