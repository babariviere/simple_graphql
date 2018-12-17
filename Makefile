define USAGE
build: build and install servers/client in go/bin
generate: generate go files from graphql and proto
endef

all: dependencies generate build

dependencies:
	go get exo/...
	go get exo_server/...

build:
	go build exo/...
	go build exo_server/...
	go install exo_server/...

generate:
	go generate exo_server/...

help:
	@echo "$(USAGE)"

run: build
	@./go/bin/server &
	@./go/bin/project &
	@./go/bin/user &

graphql: build
	@./go/bin/server

project: build
	@./go/bin/project

user: build
	@./go/bin/user
