BUILD := $(shell git rev-parse --short HEAD)
NAME := $(shell basename "$(PWD)")
VERSION := $(shell git describe --tags)

LDFLAGS=-ldflags "-s -w -X=main.build=$(BUILD) -X=main.name=$(NAME) -X=main.version=$(VERSION)" 

all : test compile run

benchmark :
	@ab -n 1000 -c 1000 http://localhost:3000/

compile :
	@go build $(LDFLAGS) -o $(GOPATH)/bin/$(NAME)_api clients/api/main.go
	@go build $(LDFLAGS) -o $(GOPATH)/bin/$(NAME)_cli clients/cli/main.go
	@go build $(LDFLAGS) -o $(GOPATH)/bin/$(NAME)_web clients/web/main.go

pprof :
	@go tool pprof out.dump

profile :
	@curl -s http://localhost:3000/debug/pprof/profile?seconds=10 > out.dump

request :
	@curl -s http://localhost:3000/

run :
	@$(GOPATH)/bin/$(NAME)_api

test :
	@go mod tidy
	@go test -v ./...
