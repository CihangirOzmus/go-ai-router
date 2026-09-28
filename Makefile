BINARY_NAME=go-ai-router
BUILD_DIR=bin
CMD_PATH=./cmd/go-ai-router

.PHONY: run build test fmt vet tidy lint clean check

run:
	go run $(CMD_PATH)

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_PATH)

test:
	go test ./... -v

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

lint:
	golangci-lint run

clean:
	rm -rf $(BUILD_DIR)

check: fmt vet lint test
