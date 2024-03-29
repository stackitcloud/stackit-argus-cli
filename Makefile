SHELL=/bin/bash -e -o pipefail
PWD = $(shell pwd)

# constants
GOLANGCI_VERSION = 1.47.2
OPENAPI_GENERATOR_VERSION=v6.0.1
DOCKER_REPO = stackit-argus-cli
DOCKER_TAG = latest

.PHONY: all
all: git-hooks download tidy build ## Initializes all tools and build

out:
	mkdir -p out

.PHONY: git-hooks
git-hooks:
	git config --local core.hooksPath .githooks/

.PHONY: download
download: ## Downloads the dependencies
	go mod download

.PHONY: tidy
tidy: ## Cleans up go.mod and go.sum
	go mod tidy

.PHONY: fmt
fmt: ## Formats all code with go fmt
	go fmt ./...

.PHONY: run
run: fmt ## Run the app
	go run ./cmd/stackit-argus-cli/main.go

.PHONY: test-build
test-build: ## Tests whether the code compiles
	go build -o /dev/null ./...

.PHONY: build
build: out/bin ## Builds all binaries

GO_BUILD = mkdir -pv "$(@)" && go build -ldflags="-w -s" -o "$(@)" ./...
.PHONY: out/bin
out/bin:
	$(GO_BUILD)

GOLANGCI_LINT = bin/golangci-lint
$(GOLANGCI_LINT):
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | bash -s -- -b bin v1.50.1

.PHONY: lint
lint: fmt $(GOLANGCI_LINT) download ## Lints all code with golangci-lint
	$(GOLANGCI_LINT) run

.PHONY: lint-reports
lint-reports: out/lint.xml

.PHONY: out/lint.xml
out/lint.xml: $(GOLANGCI_LINT) out download
	$(GOLANGCI_LINT) run ./... --out-format checkstyle | tee "$(@)"

.PHONY: test
test: ## Runs all tests
	go test $(ARGS) ./cmd/stackit-argus-cli/...
	go test $(ARGS) $(shell go list ./internal/... | grep -v /argus)

.PHONY: coverage
coverage: out/report.json ## Displays coverage per func on cli
	go tool cover -func=out/cover.out

.PHONY: html-out/bin
html-coverage: out/report.json ## Displays the coverage results in the browser
	go tool cover -html=out/cover.out

.PHONY: test-reports
test-reports: out/report.json

.PHONY: out/report.json
out/report.json: out
	go test -count 1 ./... -coverprofile=out/cover.out --json | tee "$(@)"

.PHONY: clean
clean: ## Cleans up everything
	rm -rf bin out

.PHONY: docker
docker: ## Builds docker image
	docker buildx build -t $(DOCKER_REPO):$(DOCKER_TAG) .

.PHONY: ci
ci: lint-reports test-reports ## Executes lint and test and generates reports

.PHONY: clean-generate-client
clean-generate-client: ## Remove generated APIT client code
	rm -f ./internal/services/argus/api_default.go
	rm -f ./internal/services/argus/configuration.go
	rm -f ./internal/services/argus/client.go
	rm -f ./internal/services/argus/model_*.go
	rm -f ./internal/services/argus/response.go
	rm -f ./internal/services/argus/utils.go
	rm -rf ./internal/services/argus/docs

.PHONY: generate-client-code
generate-client-code: clean-generate-client ## generate API client code
	docker run --rm \
		-v ${PWD}:/local openapitools/openapi-generator-cli:${OPENAPI_GENERATOR_VERSION} generate \
		-i /local/api/ARGUS.openapi.v1.yml \
		-g go \
		--additional-properties=packageName=argus \
		-o /local/internal/services/argus

.PHONY: generate-client
generate-client: generate-client-code tidy ## genarte API client & run go mod tidy

.PHONY: help
help: ## Shows the help
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
        awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ''
