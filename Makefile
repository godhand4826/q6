NAME=$(shell go list -m)
TAG?=latest

.PHONY: build
build: doc
	go build -o $(NAME)

.PHONY: clean
clean:
	go clean -x
	$(RM) coverage.txt coverage.xml junit.xml

.PHONY: test
test: doc
	go test -v -race -cover ./...

.PHONY: report
report: bin/gotestsum bin/gocover-cobertura
	bin/gotestsum --junitfile junit.xml --format testname
	go test -covermode count -coverprofile=coverage.txt ./...
	bin/gocover-cobertura < coverage.txt | sed 's;filename=\"gox/;filename=\";g' > coverage.xml

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

bin/gotestsum:
	@echo "Installing gotestsum to ./bin"
	env GOBIN=$(CURDIR)/bin go install gotest.tools/gotestsum@v1.11.0

bin/gocover-cobertura:
	@echo "Installing gocover-cobertura to ./bin"
	env GOBIN=$(CURDIR)/bin go install github.com/t-yuki/gocover-cobertura@v0.0.0-20180217150009-aaee18c8195c

.PHONY: image
image:
	docker build --build-arg PROGRAM=$(NAME) -t $(NAME):$(TAG) .
