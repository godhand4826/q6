# q6

## Preretirement
- golang@1.22.3
- docker
- GNU Make

## Makefile
```bash
# build
make build

# test
make test

# clean
make clean

# automatically install swag to ./bin and run it
make doc

# automatically install golangci-lint to ./bin and run it
make lint

# build docker image
make image # with default tag name 'latest'
env TAG=1.0.0 make image # or provide a tag name
```

## Project structure
```bash
.
├── bin # develop tools (golangci-lint, swag)
├── docs # swag generating docs
├── lib # libraries for non-business logic
└── pkg # business logic
```

## System design
| Method                  | Time complexity |
| ----------------------- | --------------- |
| AddSinglePersonAndMatch | O(log(n))       |
| RemoveSinglePerson      | O(log(n))       |
| QuerySinglePeople       | O(q * log(n))   |
