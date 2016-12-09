package lc

import(
  "os"
  "fmt"
  "bytes"
  "github.com/najeira/ltsv"
  "github.com/codegangsta/cli"
)

func Filter( context *cli.Context ) {
  keys, err := Keys( context )

  if err != nil {
    fmt.Fprintln( os.Stderr, err )
    os.Exit( 1 )
  }

  Scan( func( _ string, record map[ string ]string ) {
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
