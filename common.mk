VERSION=$(shell git describe --tags --always)

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	go install github.com/envoyproxy/protoc-gen-validate@v1.0.1
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/vektra/mockery/v2@v2.20.0
	go install ariga.io/atlas/cmd/atlas@v0.12.0
	go install github.com/bufbuild/buf/cmd/buf@v1.10.0

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

.PHONY: check-goreleaser-tool
check-atlas-tool:
	@if ! command -v atlas >/dev/null 2>&1; then \
		echo "altas is not installed. Please run \"make init\" or install the tool manually."; \
		exit 1; \
	fi

.PHONY: check-wire-tool
check-wire-tool:
	@if ! command -v wire >/dev/null 2>&1; then \
		echo "wire is not installed. Please run \"make init\" or install the tool manually."; \
		exit 1; \
	fi

.PHONY: check-golangci-lint-tool
check-golangci-lint-tool:
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		echo "golangci-lint is not installed. Please run \"make init\" or install the tool manually."; \
		exit 1; \
	fi

.PHONY: check-buf-tool
check-buf-tool:
	@if ! command -v buf >/dev/null 2>&1; then \
		echo "buf is not installed. Please run \"make init\" or install the tool manually."; \
		exit 1; \
	fi

