package lc

import(
  "strings"
  "errors"
  "regexp"
  "github.com/codegangsta/cli"
)

func Word( context *cli.Context ) (*regexp.Regexp, error) {
  if len(context.Args()) == 0 {
    return nil, errors.New( "Word not given" )
  }

  word := context.Args()[ 0 ]

  if len(word) <= 0 {
    return nil, errors.New( "Word not given" )
  } else {
    return regexp.MustCompile( word ), nil
  }
}

func Keys( context *cli.Context ) (map[ string ]bool, error) {
  arg := context.String( "key" )
  mapped := make( map[ string ]bool )
  splitted := strings.SplitN( arg, ",", 0 )

  for _, key := range splitted {
    mapped[ key ] = true
  }

  if len( mapped ) > 0 {
    return mapped, nil
  } else {
    return mapped, errors.New( "Key(s) not given" )
  }
}
