package main

import(
  "os"
  "fmt"
  "bytes"
  "github.com/codegangsta/cli"
  "github.com/najeira/ltsv"
  "lc"
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
  keys, _ := lc.Keys( context )
  pattern, err := lc.Word( context )

  if err != nil {
    fmt.Fprintln( os.Stderr, err )
    os.Exit( 1 )
  }

  lc.Scan( func( line string, record map[ string ]string ) {
    if len( keys ) > 0 {
      FilteringWithKey:
        for field, value := range record {
          if keys[ field ] && pattern.MatchString( value ) {
            fmt.Println( line )
            break FilteringWithKey
          }
        }
    } else {
      FilteringWithoutKey:
        for _, value := range record {
          if pattern.MatchString( value ) {
            fmt.Println( line )
            break FilteringWithoutKey
          }
        }
    }
  })
}

func filter( context *cli.Context ) {
  keys, err := lc.Keys( context )

  fmt.Printf("%d, %v\n", len(keys), keys)

  if err != nil {
    fmt.Fprintln( os.Stderr, err )
    os.Exit( 1 )
  }

  lc.Scan( func( _ string, record map[ string ]string ) {
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

func slice( record map[ string ]string, keys map[ string ]bool ) map[ string ]string {
  sliced := make( map[ string ]string )

  for key, value := range record {
    if keys[ key ] {
      sliced[ key ] = value
    }
  }

  return sliced
}
