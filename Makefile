NAME=$(shell go list -m)
TAG?=latest
LD_FLAGS=

.PHONY: build
build:
	go build -o $(NAME) .

.PHONY: clean
clean:
	go clean

.PHONY: test
test:
	go test -v -race -cover ./...

.PHONY: lint
lint: bin/golangci-lint
	bin/golangci-lint run -v

bin/golangci-lint:
	@echo "Installing golangci-lint to ./bin"
	env GOBIN=$(CURDIR)/bin go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.58.1

.PHONY: image
image:
	docker build --build-arg PROGRAM=$(NAME) --build-arg LD_FLAGS="$(LD_FLAGS)" -t $(NAME):$(TAG) .