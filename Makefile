BIN_NAME=apiseedproject

GOFILES=$(shell find . -type f -regex '^\./vendor/.*' -prune -o -name '*.go' -print)
GOPACKAGES=$(shell scripts/list-go-packages.sh)

.PHONY: clean

all: build

build: $(GOFILES)
	go build -o ./bin/$(BIN_NAME)

test: $(GOFILES)
	go test -cover $(GOPACKAGES)

run: build
	./bin/$(BIN_NAME)

clean:
	rm -f ./bin/$(BIN_NAME)
