package lc

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/sugilog/ltsv-cli.go/formatter"
	"github.com/sugilog/ltsv-cli.go/io"
)

// no test
func Filter(context *cli.Context) {
	keys, err := Keys(context)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	filter := formatter.FilterFormatter(keys)
	iomap := io.IOMap{
		Out: os.Stdout,
		Err: os.Stderr,
		In:  os.Stdin,
	}
	io.Worker(iomap, filter)
}

// no test
func Grep(context *cli.Context) {
	keys, _ := Keys(context)
	pattern, err := Word(context)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	grep := formatter.GrepFormatter(pattern, keys)
	iomap := io.IOMap{
		Out: os.Stdout,
		Err: os.Stderr,
		In:  os.Stdin,
	}
	io.Worker(iomap, grep)
}

func Word(context *cli.Context) (*regexp.Regexp, error) {
	if len(context.Args()) == 0 {
		return nil, errors.New("Word not given")
	}

	word := context.Args()[0]

	if len(word) <= 0 {
		return nil, errors.New("Word not given")
	} else {
		return regexp.MustCompile(word), nil
	}
}

func Keys(context *cli.Context) ([]string, error) {
	arg := context.String("key")
	splitted := strings.Split(arg, ",")

	if len(splitted) == 1 && splitted[0] == "" {
		return splitted, errors.New("Key(s) not given")
	} else {
		return splitted, nil
	}
}
