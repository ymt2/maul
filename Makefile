export GO111MODULE = on

PACKAGES ?= $(shell go list ./...)

GO_TEST_TARGET ?= .

REVIEWDOG_ARG ?= -diff="git diff master"

LINT_TOOLS=$(shell cat tool/tool.go | egrep '^\s_ ' | awk '{ print $$2 }')

.PHONY: all
all: test

.PHONY: bootstrap-lint-tools
bootstrap-lint-tool: # Install/Update tools
	@echo "Installing/Updating tools (dir: $(GOBIN), tools: $(LINT_TOOLS))"
	@go install -tags tool -mod=readonly $(LINT_TOOLS)

.PHONY: run
run: ## Run maul
	@go run cmd/maul/*.go ${ARGS}

.PHONY: test
test:  ## Run go test
	@go test -v -race -mod=readonly -run=$(GO_TEST_TARGET) $(PACKAGES)

.PHONY: reviewdog
reviewdog: bootstrap-lint-tools  ## Run reviewdog
	reviewdog -conf=.reviewdog.yml $(REVIEWDOG_ARG)

.PHONY: help
help:  ## Show this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[33m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z\/_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
