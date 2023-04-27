GOLANGCI_VERSION:=1.52.2
PROJECT_NAME:=merkle-tree
GOPATH_BIN:=$(shell go env GOPATH)/bin

.PHONY: install
install:
	# Install golangci-lint for go code linting.
	curl -sSfL \
		"https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh" | \
		sh -s -- -b ${GOPATH_BIN} v${GOLANGCI_VERSION}

.PHONY: all
all: lint test

.PHONY: lint
lint:
	@echo ">>> Performing golang code linting.."
	golangci-lint run --config=.golangci.yml

.PHONY: test
test:
	@echo ">>> Running Unit Tests..."
	go test -v -race ./...

.PHONY: cover-test
cover-test:
	@echo ">>> Running Tests with Coverage..."
	go test -v -race ./... -coverprofile=coverage.txt -covermode=atomic
