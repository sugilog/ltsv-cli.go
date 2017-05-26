PROGNAME = lc
SOURCE   = cmd/${PROGNAME}/main.go
GENDIR   = pkg

deps:
	go get github.com/codegangsta/cli
	go get github.com/najeira/ltsv

generate:
	env GOOS=linux  GOARCH=amd64 go build -ldfrag="-s -w" -o ${GENDIR}/${PROGNAME}.linux.amd64  ${SOURCE}
	env GOOS=darwin GOARCH=amd64 go build -ldfrag="-s -w" -o ${GENDIR}/${PROGNAME}.darwin.amd64 ${SOURCE}

sample:
	go run bin/sampler.go --line 100000 > sample.ltsv
