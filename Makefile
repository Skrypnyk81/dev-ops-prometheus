APP := $(shell basename $(shell git remote get-url origin))
REGISTRY := skrypnyk81
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGETOS=darwin
TARGETARCH=arm64

format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v

get:
	go get

build: format get
	CGO_ENABLE=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o kbot -ldflags "-X="github.com/Skrypnyk81/Dev-ops-prometheus/cmd.appVersion=${VERSION}

image:
	docker build . -t ${REGISTRY}/$(APP):${VERSION}-${TARGETARCH}

clean:
	rm -rf kbot

