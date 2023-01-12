build:
	go build .

test:
	go test ./config/... ./pkg/...

integration-test:
	go test ./integration_test/integration_test.go

test-all: test integration-test

lint:
	golangci-lint run -v

clean:
	rm -rf stackit-argus-cli

.PHONY: build test integration-test test-all lint install-test-dependencies
