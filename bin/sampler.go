package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

var src = rand.NewSource(time.Now().UnixNano())
var now = time.Now().Unix()

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)

	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func RandTime() string {
	return time.Unix(rand.Int63n(now), 0).Format(time.RFC3339)
}

func main() {
	var num int
	flag.IntVar(&num, "line", 1, "number of lines")
	flag.Parse()

	for i := 0; i < num; i++ {
		fmt.Printf("index:%d\ttext:%s\ttime:%s\n", i, RandStringBytesMaskImprSrc(8), RandTime())
	}
}
