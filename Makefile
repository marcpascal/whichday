.DEFAULT_GOAL = all

version  := $(shell git rev-parse --short HEAD)

name     := whichday
package  := github.com/marcpascal/$(name)
packages := $(shell go list ./... | grep -v /vendor/)

.PHONY: all
all:: build

.PHONY: build
build::
	go build -o whichday .

.PHONY: test
test::
	go test -v $(packages)

.PHONY: bench
bench::
	go test -bench=. -v $(packages)

.PHONY: lint
lint::
	go vet -v $(packages)

.PHONY: check
check:: lint test

.PHONY: clean
clean::
	rm -f whichday
