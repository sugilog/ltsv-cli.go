package lc

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func Grep(context *cli.Context) {
	keys, _ := Keys(context)
	pattern, err := Word(context)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	Scan(func(line string, record map[string]string) {
		if len(keys) > 0 {
		FilteringWithKey:
			for field, value := range record {
				if keys[field] && pattern.MatchString(value) {
					fmt.Println(line)
					break FilteringWithKey
				}
			}
		} else {
		FilteringWithoutKey:
			for _, value := range record {
				if pattern.MatchString(value) {
					fmt.Println(line)
					break FilteringWithoutKey
				}
			}
		}
	})
}
