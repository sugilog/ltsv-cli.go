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

func grep( context *cli.Context ) {
  var keys map[string]bool

  if len( context.String( "key" ) ) > 0  {
    keys = mapKeys( context.String( "key" ) )
  }

  word := context.Args()[ 0 ]

  scanner := bufio.NewScanner( os.Stdin )

  for scanner.Scan() {
    line   := scanner.Text()
    bytes  := bytes.NewBufferString( line )
    reader := ltsv.NewReader( bytes )
    record, err := reader.Read()

    if err != nil {
      fmt.Fprintln( os.Stderr, "parsing ltsv:", err )
    }

    if len( keys ) > 0 {
      FilteringWithKey:
        for field, value := range record {
          if keys[ field ] && value == word {
            fmt.Println( line )
            break FilteringWithKey
          }
        }
    } else {
      FilteringWithoutKey:
        for _, value := range record {
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

// func slice( record map[ string ]string, keys map[ string ]bool ) ( sliced map[ string ]string ) {
//   for key, value := range record {
//     if keys[ key ] {
//       sliced[ key ] = value
//     }
//   }
// }

func mapKeys( keys string ) ( mapped map[ string ]bool ) {
  mapped = make( map[ string ]bool )
  splitted := strings.Split( keys, "," )

  for _, key := range splitted {
    mapped[ key ] = true
  }

  return mapped
}
