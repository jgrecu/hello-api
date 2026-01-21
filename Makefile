GO_VERSION := 1.25
TAG := $(shell git describe --abbrev=0 --tags --always)
HASH := $(shell git rev-parse --short HEAD)

DATE := $(shell date +%Y-%m-%d.%H:%M:%S)
LDFLAGS := -w -X github.com/jgrecu/hello-api/handlers.hash=$(HASH) \
			  -X github.com/jgrecu/hello-api/handlers.tag=$(TAG) \
			  -X github.com/jgrecu/hello-api/handlers.date=$(DATE)

.PHONY: install-go init-go build

setup: install-go init-go install-lint copy-hooks


# TODO add MacOS support
install-go:
	wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz"
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	rm go$(GO_VERSION).linux-amd64.tar.gz

init-go:
	echo 'export PATH=$$PATH:/usr/local/go/bin' >> $${HOME}/.bashrc
	echo 'export PATH=$$PATH:$${HOME}/go/bin' >> $${HOME}/.bashrc

# TODO add MacOS support
upgrade-go:
	sudo rm -rf /usr/bin/go
	wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz"
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	rm go$(GO_VERSION).linux-amd64.tar.gz

install-lint:
	sudo curl -sSfL \
     https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh\
     | sh -s -- -b $$(go env GOPATH)/bin v1.41.1

build:
	go build -ldflags "$(LDFLAGS)" -o api cmd/main.go

test:
	go test ./... -coverprofile=coverage.out

coverage:
	go tool cover -func coverage.out \
	| grep "total:" | awk '{print ((int($$3) > 80) != 1) }'

report:
	go tool cover -html=coverage.out -o cover.html

check-format:
	test -z $(go fmt ./...)

static-check:
	golangci-lint run

clean:
	rm api coverage.out cover.html

copy-hooks:
	chmod +x scripts/hooks/*
	cp -r scripts/hooks .git/.
