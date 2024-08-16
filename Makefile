SHELL=/bin/bash -o pipefail
.DEFAULT_GOAL := help

PKG := github.com/sawadashota/tesla-home-powerflow-optimizer

GOBIN := $(abspath bin)
export PATH := $(GOBIN):${PATH}

# Generate sources from the OpenAPI spec
OAPI_CODEGEN := go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
API_SPEC := ./docs/openapi.yaml

.PHONY: tsp
tsp: ## Generate tsp
	npm run tsp
	cp -r ./tsp-output/@typespec/openapi3/openapi.yaml ./docs/openapi.yaml

.PHONY: fmt
fmt: ## Format sources
	go run golang.org/x/tools/cmd/goimports@latest -local $(PKG) -w .

.PHONY: authenticate
authenticate: ## Sign in with Tesla
	go run . authenticate

.PHONY: dev
dev: ## Run dev server
	go run . serve

.PHONY: unit-test
unit-test: ## unit test
	go test -failfast -count=1 -race ./...

.PHONY: build-html ## Build the HTML interface
build-html:
	npm --prefix interfaces/html run build

.PHONY: build ## Build the application
build: build-html
	go install .

.PhONY: generate-sdk
generate-sdk: tsp ## Generate SDK
	$(OAPI_CODEGEN) -config .oapi-codegen.yaml $(API_SPEC)
	make fmt

.PHONY: generate-ent
generate-ent: ## Generate ent code
	go generate ./ent
	make fmt

# https://gist.github.com/tadashi-aikawa/da73d277a3c1ec6767ed48d1335900f3
.PHONY: $(shell grep -h -E '^[a-zA-Z_-]+:' $(MAKEFILE_LIST) | sed 's/://')

# https://postd.cc/auto-documented-makefile/
help: ## Show help message
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
