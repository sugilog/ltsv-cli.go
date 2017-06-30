PROGNAME = lc
SOURCE   = cmd/${PROGNAME}/main.go
GENDIR   = pkg

.DEFAULT_GOAL := help

## Run tests.
test:
	go test ./ ./io ./formatter

## Install dependencies
deps:
	go get -u github.com/codegangsta/cli
	go get -u github.com/najeira/ltsv

## Build releases
release:
	env GOOS=linux  GOARCH=amd64 go build -ldfrag="-s -w" -o ${GENDIR}/${PROGNAME}.linux.amd64  ${SOURCE}
	env GOOS=darwin GOARCH=amd64 go build -ldfrag="-s -w" -o ${GENDIR}/${PROGNAME}.darwin.amd64 ${SOURCE}

## Generate sample.ltsv
sample:
	go run bin/sampler.go --line 100000 > sample.ltsv

## Show help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: test deps release sample
