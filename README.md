# q6

## Dependency
- golang@1.22.3
- docker
- GNU Make

## Build
```bash
make build
```

## Test
```bash
make test
```

## Clean
```bash
make clean
```

## Lint
> The script will automatically install `golangci-lint` to `./bin`
```bash
make lint
```

## Docker Image
```bash
# use the default tag name 'latest'
make image

# or provide a tag name
env TAG=1.0.0 make image
```