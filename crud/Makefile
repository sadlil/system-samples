# It's necessary to set this because some environments don't link sh -> bash.
#SHELL := /bin/bash

DIRs := $(shell (go list ./... | grep -v /vendor/))

.PHONY: all
all: gen fmt test install

.PHONY: gen
gen: gen-proto-go
	go-bindata -o apis/openapi/specs.go -pkg openapi apis/openapi/gen
	go generate ./...

.PHONY: install
install: gen
	go install ./...

.PHONY: build
build: install

.PHONY: fmt
fmt: gen
	@./hack/scripts/run_fmt.sh

lint:
	buf lint apis/
	@golangci-lint run --out-format code-climate | \
		tee gl-code-quality-report.json | \
		jq -r '.[] | "\(.location.path):\(.location.lines.begin) \(.description)"'
	@rm gl-code-quality-report.json

.PHONE: dep
dep: dep-proto
	go mod download
	go mod tidy -e

.PHONY: dep.proto
dep-proto:
	buf mod update apis/

gen-proto-go:
	buf generate apis/ --template=apis/buf.crudapi.yaml
	mockery

# Build related commands
build.docker:
	docker build -t github.com/sadlil/system-samples/crud:$(VERSION) -f hack/docker/Dockerfile ./..

TEST_TIMEOUT = 60
TEST_PKGS ?= $(shell (go list ./... | grep -v /vendor/))
TEST_TARGETS := test-short test-verbose test-race test-cover test-cover-report
.PHONY: $(TEST_TARGETS) test tests
test-short:   TEST_ARGS=-short
test-verbose: TEST_ARGS=-v
test-race:    TEST_ARGS=-race
test-cover:   TEST_ARGS=-cover -race
test-cover-report:  TEST_ARGS=-cover -race -coverprofile=coverage.txt -covermode=atomic
$(TEST_TARGETS): test
test: gen
	@go test -timeout $(TEST_TIMEOUT)s $(TEST_ARGS) $(TEST_PKGS)

clean:
	go clean -i -r -x
	rm gl-code-quality-report.json
