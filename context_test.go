package lc

import (
	"flag"
	"testing"

	"github.com/codegangsta/cli"
)

// func TestWord(t *testing.T) {
// 	set := flag.NewFlagSet("test", 0)
// 	c := cli.NewContext(nil, set, nil)

// 	set.Parse([]string{"grep word", "--key", "field1"})

// 	pattern, err := Word(c)

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	if !pattern.MatchString("Is there grep word ??") {
// 		t.Error("Word should compile 'grep word'", pattern)
// 	}

// 	if pattern.MatchString("Is there rep word ??") {
// 		t.Error("Word should compile 'grep word'", pattern)
// 	}
// }

// func TestWordWithEmptyArgument(t *testing.T) {
// 	set := flag.NewFlagSet("test", 0)
// 	c := cli.NewContext(nil, set, nil)

// 	var empty []string
// 	set.Parse(empty)

// 	_, err := Word(c)

// 	if err == nil {
// 		t.Error("Should return error")
// 	}

// 	if err.Error() != "Word not given" {
// 		t.Error("Error Message should say 'Word not given'")
// 	}
// }

func TestKeys(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	set.String("key", "", "doc")
	c := cli.NewContext(nil, set, nil)

	set.Parse([]string{"--key", "field1"})

	keys, err := Keys(c)

	if err != nil {
		t.Error(err)
	}

	if !keys["field1"] {
		t.Error("Keys should return with key 'field1'")
	}
}

func TestKeysHavingComma(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	set.String("key", "", "doc")
	c := cli.NewContext(nil, set, nil)

	set.Parse([]string{"--key", "field1,field2"})

	keys, err := Keys(c)

	if err != nil {
		t.Error(err)
	}

	if !keys["field1"] {
		t.Error("Keys should return with key 'field1'")
	}

	if !keys["field2"] {
		t.Error("Keys should return with key 'field2'")
	}
}

func TestKeysWithEmptyArgs(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	set.String("key", "", "doc")
	c := cli.NewContext(nil, set, nil)

	set.Parse([]string{"--key", ""})

	_, err := Keys(c)

	if err == nil {
		t.Error("Should return error")
	}
}

func TestKeysWithNoFlag(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	c := cli.NewContext(nil, set, nil)

	_, err := Keys(c)

	if err == nil {
		t.Error("Should return error")
	}
}
