# It's necessary to set this because some environments don't link sh -> bash.
#SHELL := /bin/bash

DIRs := $(shell (go list ./... | grep -v /vendor/))

.PHONY: all
all: gen fmt test install

.PHONY: gen
gen: gen-proto-go
	go-bindata -o apis/openapi/specs.go -pkg openapi apis/openapi/apis
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
	@golangci-lint run --out-format code-climate | \
		tee gl-code-quality-report.json | \
		jq -r '.[] | "\(.location.path):\(.location.lines.begin) \(.description)"'
	@rm gl-code-quality-report.json

.PHONE: dep
dep:
	go mod download
	go mod tidy -e

.PHONY: dep.proto
dep-proto:
	protodep up -u -f

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

gen-proto-go:
	@protoc -I . \
	   -I third_party/protos/googleapis \
	   -I third_party/protos/envoyproxy \
	   -I third_party/protos \
       --go_out ./apis/go \
       --go-grpc_out ./apis/go \
       --validate_out="lang=go:./apis/go" \
       --grpc-gateway_out ./apis/go \
       --grpc-gateway_opt logtostderr=true \
       --grpc-gateway_opt generate_unbound_methods=true \
       --openapiv2_out ./apis/openapi/apis \
       --openapiv2_opt logtostderr=true,allow_merge=true,use_go_templates=true \
       apis/*.proto
	@mv apis/openapi/apis/apidocs.swagger.json apis/openapi/apis/swagger.json

clean:
	go clean -i -r -x
	rm gl-code-quality-report.json


# Build related commands
build.docker:
	docker build -t sadlil.com/samples/crud:$(VERSION) -f hack/docker/Dockerfile ./..