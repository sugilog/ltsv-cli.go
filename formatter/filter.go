package formatter

import (
	"strings"

	"github.com/sugilog/ltsv-cli.go/io"
)

func FilterFormatter(keys []string) func(entry io.Entry) (io.Entry, bool) {
	return func(entry io.Entry) (io.Entry, bool) {
		ltsv, err := entry.LTSV()

		if err != nil {
			return entry, false
		}

		entry.Formatted = FilterFunc(ltsv, keys)

		return entry, true
	}
}

func FilterFunc(ltsv map[string]string, keys []string) string {
	arr := make([]string, 0, len(ltsv))

	for _, key := range keys {
		if value, ok := ltsv[key]; ok {
			arr = append(arr, key+":"+value)
		}
	}

	return strings.Join(arr, "\t")
}
