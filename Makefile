BIN_NAME=apiseedproject

GOFILES=$(shell find . -type f -regex '^\./vendor/.*' -prune -o -name '*.go' -print)
GOPACKAGES=$(shell go list ./... | egrep -v 'github.com/apiseedprojects/go/vendor/')

.PHONY: clean

all: build

build: $(GOFILES)
	go build -o ./bin/$(BIN_NAME)

test: $(GOFILES)
	go test -cover -v $(GOPACKAGES)

run: build
	./bin/$(BIN_NAME)

clean:
	rm -f ./bin/$(BIN_NAME)
