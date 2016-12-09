PROGNAME="lc"
SOURCE="main.go"
GENDIR="gen"

gen:
	env GOOS=linux  GOARCH=386   gom build -o ${GENDIR}/${PROGNAME}.linux.386    ${SOURCE}
	env GOOS=linux  GOARCH=arm   gom build -o ${GENDIR}/${PROGNAME}.linux.arm    ${SOURCE}
	env GOOS=linux  GOARCH=amd64 gom build -o ${GENDIR}/${PROGNAME}.linux.amd64  ${SOURCE}
	env GOOS=darwin GOARCH=386   gom build -o ${GENDIR}/${PROGNAME}.darwin.386   ${SOURCE}
	env GOOS=darwin GOARCH=amd64 gom build -o ${GENDIR}/${PROGNAME}.darwin.amd64 ${SOURCE}
sample:
	go run bin/sampler.go --line 100000 > sample.ltsv
