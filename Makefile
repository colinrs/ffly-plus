SHELL := /bin/bash
BASEDIR = $(shell pwd)


# build with verison infos
versionDir = "ffly-plus/internal/version"
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

ldflags="-w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState}"

PROJECT_NAME := "ffly-plus"
PKG := "$(PROJECT_NAME)"

PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all
all: build

.PHONY: build
build: ## Build the binary file
	@sh check.sh
	@protoc -I ./internal internal/proto/*.proto --go_out=plugins=grpc:./internal
	@swag init
	@go build -v -ldflags ${ldflags} .
	@gofmt -w .


.PHONY: clean
clean:
	rm -f ffly-plus
	rm cover.out coverage.txt
	find . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {}

.PHONY: test
test: ## view test result
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST}
	@cat cover.out > coverage.txt
	@go tool cover -html=coverage.txt

.PHONY: docs
docs:
	@swag init
	@echo "gen-docs done"
	@echo "see docs by: http://localhost:8080/swagger/index.html"

.PHONY: ca
ca:
	openssl req -new -nodes -x509 -out config/server.crt -keyout config/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"

.PHONY: proto
proto:
	protoc -I ./internal internal/proto/*.proto --go_out=plugins=grpc:./internal

.PHONY: help
help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make ca - generate ca files"
	@echo "make docs - gen swag doc"
	@echo "make test - go test"
	@echo "make build - go build"
	@echo "make proto - build proto"


