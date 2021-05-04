BUILD := $(shell git rev-parse --short HEAD)
NAME := $(shell basename "$(PWD)")
VERSION := $(shell git describe --tags)

LDFLAGS=-ldflags "-s -w -X=main.build=$(BUILD) -X=main.name=$(NAME) -X=main.version=$(VERSION)" 

all : test compile compress run

benchmark :
	@ab -n 1000 -c 1000 http://localhost:3000/

compile :
	@go build $(LDFLAGS) -o build/$(NAME)_api clients/api/main.go

compress :
	@upx -9 build/$(NAME)_api

pprof :
	@go tool pprof out.dump

profile :
	@curl -s http://localhost:3000/debug/pprof/profile?seconds=10 > out.dump

request :
	@curl -s http://localhost:3000/

run :
	@./build/$(NAME)_api

test :
	@go test -v ./...
