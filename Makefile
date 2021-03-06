.PHONY: default build image check publish-images

TAG_NAME := $(shell git tag -l --contains HEAD)

default: check test build

test:
	go test -v -cover ./...

build:
	CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o kws_sig

dep:
	CGO_ENABLED=0 go mod download
build_bin:
	CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o bin/sig

image:
	docker build -t kws_sig .

check:
	golangci-lint run

publish-images:
	seihon
# 	publish -v "$(TAG_NAME)" -v "latest" --image-name containous/whoami --dry-run=false