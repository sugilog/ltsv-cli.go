package main

import(
  "os"
  "lc"
  "github.com/codegangsta/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "lc"
  app.Usage = "ltsv grep/filtering tool"
  app.Commands = []cli.Command {
    {
      Name:    "grep",
      Aliases: []string{ "g" },
      Usage:   "grep by specified word in specfied keys",
      Action:  lc.Grep,
      Flags:   []cli.Flag {
        cli.StringFlag {
          Name: "key, k",
          Value: "",
          Usage: "greppable key",
        },
      },
    },
    {
      Name:    "filter",
      Aliases: []string{ "f" },
      Usage:   "filter ltsv fieds only specified keys",
      Action:  lc.Filter,
      Flags:   []cli.Flag {
        cli.StringFlag {
          Name: "key, k",
          Value: "",
          Usage: "slice key",
        },
      },
    },
  }

  app.Run(os.Args)
}
