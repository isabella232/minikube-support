SHELL = /usr/bin/env bash -o pipefail -o errexit -o nounset
NAME := minikube-support
ORG := chr-fritz
ROOT_PACKAGE := github.com/qaware/minikube-support
VERSION := v0.0.0-next

REVISION   := $(shell git rev-parse --short HEAD 2> /dev/null  || echo 'unknown')
BRANCH     := $(shell git rev-parse --abbrev-ref HEAD 2> /dev/null  || echo 'unknown')
BUILD_DATE := $(shell git show -s --format=%ct)

GO_VERSION=$(shell go version | sed -e 's/^[^0-9.]*\([0-9.]*\).*/\1/')
PACKAGE_DIRS := $(shell go list ./...)

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
BUILD_DIR ?= ./bin
REPORTS_DIR ?= ./reports

BUILDFLAGS := -ldflags \
  " -X '$(ROOT_PACKAGE)/version.Version=$(VERSION)'\
    -X '$(ROOT_PACKAGE)/version.Revision=$(REVISION)'\
    -X '$(ROOT_PACKAGE)/version.Branch=$(BRANCH)'\
    -X '$(ROOT_PACKAGE)/version.CommitDate=$(BUILD_DATE)'\
    -s -w -extldflags '-static'"

.PHONY: all
all: lint test $(GOOS)-build
	@echo "SUCCESS"

.PHONY: ci
ci: ci-check

.PHONY: ci-check
ci-check: lint tidy generate imports vet
	git diff --exit-code

check: fmt test

.PHONY: build
build: pb
	CGO_ENABLED=0 GOARCH=amd64 go build $(BUILDFLAGS) -o $(BUILD_DIR)/$(NAME) $(ROOT_PACKAGE)

.PHONY: debug
debug: pb
	CGO_ENABLED=0 GOARCH=amd64 go build -gcflags "all=-N -l" -o $(BUILD_DIR)/$(NAME)-debug $(ROOT_PACKAGE)
	dlv --listen=:2345 --headless=true --api-version=2 exec $(BUILD_DIR)/$(NAME)-debug run

.PHONY: imports
imports:
	find . -type f -name '*.go' ! -name '*_mocks.go' -print0 | xargs -0 goimports -w -l

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: darwin-build
darwin-build: pb
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build $(BUILDFLAGS) -o $(BUILD_DIR)/$(NAME)-darwin $(ROOT_PACKAGE)

.PHONY: linux-build
linux-build: pb
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build $(BUILDFLAGS) -o $(BUILD_DIR)/$(NAME)-linux $(ROOT_PACKAGE)

.PHONY: windows-build
windows-build: pb
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build $(BUILDFLAGS) -o $(BUILD_DIR)/$(NAME)-windows.exe $(ROOT_PACKAGE)

.PHONY: test
test: generate pb
	mkdir -p $(REPORTS_DIR)
	go test $(PACKAGE_DIRS) -coverprofile=$(REPORTS_DIR)/coverage.out -v $(PACKAGE_DIRS) | tee >(go tool test2json > $(REPORTS_DIR)/tests.json)

.PHONY: test-race
test-race: generate pb
	mkdir -p $(REPORTS_DIR)
	go test -race $(PACKAGE_DIRS) -coverprofile=$(REPORTS_DIR)/coverage.out -v $(PACKAGE_DIRS) | tee >(go tool test2json > $(REPORTS_DIR)/tests.json)

.PHONY: cross
cross: darwin-build linux-build windows-build

.PHONY: pb
pb:
	$(MAKE) -C pb

.PHONY: vet
vet:
	mkdir -p $(REPORTS_DIR)
	go vet -v $(PACKAGE_DIRS) 2> >(tee $(REPORTS_DIR)/vet.out) || true

.PHONY: lint
lint:
	mkdir -p $(REPORTS_DIR)
	# GOGC default is 100, but we need more aggressive GC to not consume too much memory
	# might not be necessary in future versions of golangci-lint
	# https://github.com/golangci/golangci-lint/issues/483
	GOGC=20 golangci-lint run --disable=typecheck --deadline=5m --out-format checkstyle > $(REPORTS_DIR)/lint.xml || true

.PHONY: generate
generate:
	go generate ./...

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
	rm -rf release
	rm -rf $(REPORTS_DIR)

.PHONY: buildDeps
buildDeps:
	go mod download
	go get -u google.golang.org/grpc
	go get -u github.com/golang/protobuf/protoc-gen-go
	go install github.com/golang/mock/mockgen@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	$(MAKE) -C pb buildDeps

.PHONY: completions
completions:
	rm -rf completions
	mkdir completions
	for sh in bash zsh fish ps1; do go run main.go completion "$$sh" >"completions/$(NAME).$$sh"; done
