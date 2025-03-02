# It's necessary to set this because some environments don't link sh -> bash.
#SHELL := /bin/bash

DIRs := $(shell (go list ./... | grep -v /vendor/))

.PHONY: all
all: install

.PHONY: install
install:
	go install ./...

.PHONY: build
build: install

lint:
	buf lint apis/
	@golangci-lint run --out-format code-climate | \
		tee gl-code-quality-report.json | \
		jq -r '.[] | "\(.location.path):\(.location.lines.begin) \(.description)"'
	@rm gl-code-quality-report.json

.PHONE: dep
dep:
	go mod download
	go mod tidy -e

# Build related commands
build.docker:
	docker build -t github.com/sadlil/system-samples/periscope:$(VERSION) -f Dockerfile ./..

clean:
	go clean -i -r -x
	rm gl-code-quality-report.json
