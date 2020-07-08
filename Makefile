.PHONY: help default build lint push

NAME ?= aaneci/grpc-sample-app
OS ?= linux
COMMIT ?= $(shell git rev-parse --short HEAD)


default: help

help: ## Print this help text
	@printf "\n"
	@awk 'BEGIN {FS = ":.*?## "}; ($$2 && !/@awk/){printf "${CYAN}%-30s${NC} %s\n", $$1, $$2}' $(lastword ${MAKEFILE_LIST}) | sort
	@printf "\n"

# CI
# -----------
lint: ## Format and lint go code
	test -z "$(gofmt -l ./...)"
	test -z "$(golint ./...)"

build: lint ## Builds the local Docker container for development
	docker build --build-arg OS=${OS} -t ${NAME}:$(COMMIT) .
	docker tag ${NAME}:$(COMMIT) ${NAME}:latest

push: build ## Push Docker image to artifactory
	docker push ${NAME}:$(COMMIT)
	docker push ${NAME}:latest
