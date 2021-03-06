SHELL := /bin/bash

GIT_SHA := $(shell git rev-parse HEAD | cut -c 1-12)
VERSION := $(shell git describe --tags --dirty --always --abbrev=12)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

IMAGE_NAME = zetsub0u/objcache

default: build

all:

deps:
	go mod tidy
	go mod vendor
.PHONY: deps

build: deps
	@go generate
	GOOS=linux go build -mod vendor -ldflags "-X github.com/zetsub0u/objcache/cmd.version=$(VERSION) -X github.com/zetsub0u/objcache/cmd.commit=$(GIT_SHA) -X github.com/zetsub0u/objcache/cmd.branch=$(BRANCH)" -o bin/objcache
.PHONY: build

build-darwin: deps
	@go generate
	GOOS=darwin go build -mod vendor -ldflags "-X github.com/zetsub0u/objcache/cmd.version=$(VERSION) -X github.com/zetsub0u/objcache/cmd.commit=$(GIT_SHA) -X github.com/zetsub0u/objcache/cmd.branch=$(BRANCH)" -o bin/objcache
.PHONY: build

run: deps docs build
	./bin/objcache start
.PHONY: run

run-darwin: deps docs build-darwin
	./bin/objcache start
.PHONY: run-darwin

docs:
	swag init
.PHONY: docs

docker-image: deps
	docker build --build-arg VERSION=$(VERSION) --build-arg COMMIT=$(GIT_SHA) --build-arg BRANCH=$(BRANCH) -t "$(IMAGE_NAME):$(GIT_SHA)" .
.PHONY: docker-image

docker-image-without-deps:
	docker build --build-arg VERSION=$(VERSION) --build-arg COMMIT=$(GIT_SHA) --build-arg BRANCH=$(BRANCH) -t "$(IMAGE_NAME):$(GIT_SHA)" .
.PHONY: docker-image

docker-run: docker-image
	docker run -i --rm --entrypoint "/usr/local/bin/objcache" -p 8080:8080 $(IMAGE_NAME):$(GIT_SHA) start -b 0.0.0.0
.PHONY: docker-run

test:
	@go install github.com/tebeka/go2xunit
	@rm -f gotest.out tests.xml
	@2>&1 go test -v ./... -cover -race | tee gotest.out
	@go2xunit -fail -input gotest.out -output tests.xml
.PHONY: test
