define USAGE
build: build and install servers/client in go/bin
generate: generate go files from graphql and proto
endef

GODIR=$(GOPATH)/src/exo

all: generate build test

build: FORCE
	go build ./...

generate: FORCE
	@# Hack to bypass gqlgen modules
	@if [ ! -d $(GODIR) ]; then mkdir -p $(GODIR); fi
	# We needs password because we have to bind this directory into gopath
	@mount | grep $(GODIR) &>/dev/null || sudo mount --bind `pwd` $(GODIR)
	cd $(GOPATH)/src/exo; go generate ./...
	@sudo umount $(GODIR)

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
