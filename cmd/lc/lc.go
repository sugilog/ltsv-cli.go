package main

import (
	"github.com/codegangsta/cli"
	"github.com/sugilog/ltsv-cli.go"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "lc"
	app.Usage = "ltsv grep/filtering tool"
	app.Commands = []cli.Command{
		// {
		// 	Name:    "grep",
		// 	Aliases: []string{"g"},
		// 	Usage:   "grep by specified word in specfied keys",
		// 	Action:  lc.Grep,
		// 	Flags: []cli.Flag{
		// 		cli.StringFlag{
		// 			Name:  "key, k",
		// 			Value: "",
		// 			Usage: "grep-able key",
		// 		},
		// 	},
		// },
		{
			Name:    "filter",
			Aliases: []string{"f"},
			Usage:   "filter ltsv fields only specified keys",
			Action:  lc.Filter,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "key, k",
					Value: "",
					Usage: "filtering key",
				},
			},
		},
	}

	app.Run(os.Args)
}
