NAME=$(shell go list -m)
TAG?=latest

.PHONY: build
build: doc
	go build -o $(NAME)

.PHONY: clean
clean:
	go clean -x

.PHONY: test
test: doc
	go test -v -race -cover ./...

.PHONY: lint
lint: bin/golangci-lint doc
	bin/golangci-lint run -v

.PHONY: doc
doc: bin/swag
	bin/swag init

bin/golangci-lint:
	@echo "Installing golangci-lint to ./bin"
	env GOBIN=$(CURDIR)/bin go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.58.1

bin/swag:
	@echo "Installing swag to ./bin"
	env GOBIN=$(CURDIR)/bin go install github.com/swaggo/swag/cmd/swag@v1.16.3

.PHONY: image
image:
	docker build --build-arg PROGRAM=$(NAME) -t $(NAME):$(TAG) .
