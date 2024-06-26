# Makefile

BINARY_NAME=boa

all: build test

build:
	go build -o bin/${BINARY_NAME} main.go

test: build
	bash docs/example.sh
#	go test -v main.go

run:
	go build -o bin/${BINARY_NAME} main.go
	./bin/${BINARY_NAME}

dist:
	goreleaser release --snapshot --clean

clean:
	go clean
	rm bin/${BINARY_NAME}

# https://gist.github.com/thomaspoignant/5b72d579bd5f311904d973652180c705#file-makefile


################################################################################
# general
################################################################################

pre: ## run pre-commit
	pre-commit run --all-files

help: ## show help message
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage: \033[36m\033[0m\n"} /^[$$()% 0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s:\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\asd033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
