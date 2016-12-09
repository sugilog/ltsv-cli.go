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

  if len(keys) > 0 {
    Scan(func(line string, record map[string]string) {
      for field, value := range record {
        if keys[field] && pattern.MatchString(value) {
          fmt.Println(line)
          break
        }
      }
    })
  } else {
    Scan(func(line string, record map[string]string) {
			for _, value := range record {
				if pattern.MatchString(value) {
					fmt.Println(line)
					break
				}
			}
    })
  }
}
