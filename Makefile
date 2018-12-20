define USAGE
build: build and install servers/client in go/bin
generate: generate go files from graphql and proto
endef

GOPATH=$(shell pwd)
GODIR=$(GOPATH)/src/exo

all: build test

build: build_graphql build_project build_user

build_graphql: dependencies
	 go build -o bin/graphql exo/server/graphql/server

build_project: dependencies
	 go build -o bin/project exo/server/project

build_user: dependencies
	 go build -o bin/user exo/server/user

dependencies:
	go get exo/...

generate: FORCE
	go generate exo/...

test: FORCE
	 go test exo/...

help:
	@echo "$(USAGE)"

graphql: build_graphql
	@./bin/graphql

project: build_project
	@./bin/project

user: build_user
	@./bin/user

FORCE:
