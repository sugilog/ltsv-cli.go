package lc

import(
  "os"
  "fmt"
  "bufio"
  "strings"
  "bytes"
  "github.com/najeira/ltsv"
)

func Scan( handler func( string, map[ string ]string ) ) {
  scanner := bufio.NewScanner( os.Stdin )

  for scanner.Scan() {
    line   := strings.Trim( scanner.Text(), "\t" )
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
