package formatter

import (
	"regexp"
	"sort"
	"strings"

	"github.com/sugilog/ltsv-cli.go/io"
)

func GrepFormatter(pattern *regexp.Regexp, keys []string) func(entry io.Entry) (io.Entry, bool) {
	return func(entry io.Entry) (io.Entry, bool) {
		var ok bool
		ltsv, err := entry.LTSV()

		if err != nil {
			return entry, false
		}

		if len(keys) == 0 || (len(keys) == 1 && keys[0] == "") {
			entry.Formatted, ok = GrepFuncNoKeys(ltsv, pattern)
		} else {
			entry.Formatted, ok = GrepFuncWithKeys(ltsv, pattern, keys)
		}

		return entry, ok
	}
}

func GrepFuncWithKeys(ltsv map[string]string, pattern *regexp.Regexp, keys []string) (string, bool) {
	found := false

	for _, key := range keys {
		if value, ok := ltsv[key]; ok {
			if pattern.MatchString(value) {
				found = true
				break
			}
		}
	}

	if found {
		return SortedJoin(ltsv), true
	} else {
		return "", false
	}
}

func GrepFuncNoKeys(ltsv map[string]string, pattern *regexp.Regexp) (string, bool) {
	found := false

	for _, value := range ltsv {
		if pattern.MatchString(value) {
			found = true
			break
		}
	}

	if found {
		return SortedJoin(ltsv), true
	} else {
		return "", false
	}
}

func SortedJoin(ltsv map[string]string) string {
	sorted := make([]string, len(ltsv))
	i := 0
	j := 0

	for key, _ := range ltsv {
		sorted[i] = key
		i++
	}
	sort.Strings(sorted)

	for _, key := range sorted {
		sorted[j] = key + ":" + ltsv[key]
		j++
	}

	return strings.Join(sorted, "\t")
}
