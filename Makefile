BUILD := $(shell git rev-parse --short HEAD)
NAME := $(shell basename "$(PWD)")
VERSION := $(shell git describe --tags)

LDFLAGS=-ldflags "-s -w -X=main.build=$(BUILD) -X=main.name=$(NAME) -X=main.version=$(VERSION)" 

SRC=clients/cli/main.go
BIN=build/$(NAME)

all : test compile install

compile :
	@go build $(LDFLAGS) -o $(BIN) $(SRC)

install :
	@mv $(BIN) $(GOPATH)/bin/$(NAME)

test :
	@go test -v ./...
