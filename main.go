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

func grep( context *cli.Context ) {
  var keys map[string]bool

  if len( context.String( "key" ) ) > 0  {
    keys = mapKeys( context.String( "key" ) )
  }

  word := context.Args()[ 0 ]

  scan( func( line string, record map[ string ]string ) {
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
  })
}

func filter( context *cli.Context ) {
  var keys map[string]bool

  if len( context.String( "key" ) ) > 0  {
    keys = mapKeys( context.String( "key" ) )
  } else {
    fmt.Fprintln( os.Stderr, "key should be given" )
    os.Exit( 1 )
  }

  scan( func( _ string, record map[ string ]string ) {
    sliced := slice( record, keys )

    bytes  := &bytes.Buffer{}
    writer := ltsv.NewWriter( bytes )
    // err    := writer.Write( record )
    err    := writer.WriteAll( []map[ string ]string{ sliced } )

    if err != nil {
      fmt.Fprintln( os.Stderr, "writing ltsv:", err, sliced )
    } else {
      fmt.Printf( "%v", bytes.String() )
    }
  })
}

func scan( handler func( string, map[ string ]string ) ) {
  scanner := bufio.NewScanner( os.Stdin )

  for scanner.Scan() {
    line   := scanner.Text()
    bytes  := bytes.NewBufferString( line )
    reader := ltsv.NewReader( bytes )
    record, err := reader.Read()

    if err != nil {
      fmt.Fprintln( os.Stderr, "parsing ltsv:", err )
    }

    handler( line, record )
  }

  if err := scanner.Err(); err != nil {
    fmt.Fprintln( os.Stderr, "reading standard input:", err )
  }
}

func slice( record map[ string ]string, keys map[ string ]bool ) map[ string ]string {
  sliced := make( map[ string ]string )

  for key, value := range record {
    if keys[ key ] {
      sliced[ key ] = value
    }
  }

  return sliced
}

func mapKeys( keys string ) map[ string ]bool {
  mapped := make( map[ string ]bool )
  splitted := strings.Split( keys, "," )

  for _, key := range splitted {
    mapped[ key ] = true
  }

  return mapped
}
