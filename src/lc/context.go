package lc

import(
  "strings"
  "errors"
  "github.com/codegangsta/cli"
)

func Keys( context *cli.Context ) (map[ string ]bool, error) {
  arg := context.String( "key" )
  mapped := make( map[ string ]bool )
  splitted := strings.SplitN( arg, ",", 0 )

  for _, key := range splitted {
    fmt.Println( key )
    mapped[ key ] = true
  }

  if len( mapped ) > 0 {
    return mapped, nil
  } else {
    return mapped, errors.New( "Key(s) not given" )
  }
}
