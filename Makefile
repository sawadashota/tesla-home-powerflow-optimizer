SHELL=/bin/bash -o pipefail
.DEFAULT_GOAL := help

PKG := github.com/sawadashota/tesla-home-powerflow-optimizer

GOBIN := $(abspath bin)
export PATH := $(GOBIN):${PATH}

.PHONY: fmt
fmt: ## Format sources
	go run golang.org/x/tools/cmd/goimports@latest -local $(PKG) -w .

.PHONY: unit-test
unit-test: ## unit test
	go test -failfast -count=1 -race ./...

.PHONY: ent-generate
ent-generate: ## Generate ent code
	go generate ./ent

# https://gist.github.com/tadashi-aikawa/da73d277a3c1ec6767ed48d1335900f3
.PHONY: $(shell grep -h -E '^[a-zA-Z_-]+:' $(MAKEFILE_LIST) | sed 's/://')

# https://postd.cc/auto-documented-makefile/
help: ## Show help message
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
