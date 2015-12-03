package main

import(
  "os"
  "fmt"
  "bufio"
  "bytes"
  "strings"
  "github.com/codegangsta/cli"
  "github.com/najeira/ltsv"
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
      Action:  grep,
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
      Action:  filter,
    },
  }
  app.Run(os.Args)
}

func grep( c *cli.Context ) {
  var keys []string

  if len( c.String( "key" ) ) > 0  {
    keys = strings.Split( c.String( "key" ), "," )
  }

  word := c.Args()[ 0 ]

  scanner := bufio.NewScanner( os.Stdin )

  for scanner.Scan() {
    line   := scanner.Text()
    bytes  := bytes.NewBufferString( line )
    reader := ltsv.NewReader( bytes )
    records, err := reader.ReadAll()

    if err != nil {
      fmt.Fprintln( os.Stderr, "parsing ltsv:", err )
    }

    if len( keys ) > 0 {
      FilteringWithKey:
        for field, value := range records[ 0 ] {
          for _, key := range keys {
            if field == key && value == word {
              fmt.Println( line )
              break FilteringWithKey
            }
          }
        }
    } else {
      FilteringWithoutKey:
        for _, value := range records[ 0 ] {
          if value == word {
            fmt.Println( line )
            break FilteringWithoutKey
          }
        }
    }
  }

  if err := scanner.Err(); err != nil {
    fmt.Fprintln( os.Stderr, "reading standard input:", err )
  }
}

func filter( c *cli.Context ) {
  fmt.Println( "filter" )
  fmt.Println( c )
}
