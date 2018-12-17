define USAGE
build: build and install servers/client in go/bin
generate: generate go files from graphql and proto
endef

all: generate build test

build: FORCE
	go build ./...

generate: FORCE
	go generate ./...

test: FORCE
	go test ./...

help:
	@echo "$(USAGE)"

graphql: FORCE
	@go build -o bin/graphql exo/server/graphql/server
	@./bin/graphql

project: FORCE
	@go build -o bin/project exo/server/project
	@./bin/project

user: FORCE
	@go build -o bin/user exo/server/user
	@./bin/user

FORCE:
