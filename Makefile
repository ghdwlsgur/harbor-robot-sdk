SHELL := /usr/bin/env bash

.DEFAULT_GOAL := help

DOCKERCMD=$(shell which docker)
SWAGGER_VERSION=latest
SWAGGER := $(DOCKERCMD) run --rm -it -v $(HOME):$(HOME) -w $(shell pwd) quay.io/goswagger/swagger:$(SWAGGER_VERSION)

HARBOR_SPEC=api/v2.0/swagger.yaml
HARBOR_CLIENT_DIR=pkg/sdk/robot

## --------------------------------------
## Help
## --------------------------------------

help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## --------------------------------------
## Client
## --------------------------------------

.PHONY: gen-harbor-api
gen-harbor-api: # update-spec ## generate goswagger client for Chequer
	@$(SWAGGER) generate client -f ${HARBOR_SPEC} --target=$(HARBOR_CLIENT_DIR)

.PHONY: test
test: ## run the test
	go test ./...

all: gen-harbor-api
