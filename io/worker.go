package io

import (
	"bufio"
	"fmt"
	"sort"
	"sync"
)

func Worker(iomap IOMap, handler func(Entry) (Entry, bool)) {
	outputWorker := make(chan Entry)
	outputWait := Output(&iomap, outputWorker)
	Reader(&iomap, outputWorker, handler)
	outputWait.Wait()
}

func Reader(iomap *IOMap, outputWorker chan Entry, handler func(Entry) (Entry, bool)) {
	var readWait sync.WaitGroup

	scanner := bufio.NewScanner(iomap.In)
	count := 0
	limit := 4
	var semaphore = make(chan struct{}, limit)

	defer func() {
		outputWorker <- NullEntry
	}()

	for scanner.Scan() {
		semaphore <- struct{}{}
		readWait.Add(1)
		count++

		go func(lineNo int, line string) {
			defer func() {
				readWait.Done()
				<-semaphore
			}()

			entry := Entry{LineNo: lineNo, Line: line}
			entry, ok := handler(entry)

			if ok {
				outputWorker <- entry
			}
		}(count, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(iomap.Err, "Error on reading input:", err)
	}

	readWait.Wait()
}

func Output(iomap *IOMap, outputWorker chan Entry) *sync.WaitGroup {
	var outputWait sync.WaitGroup
	outputWait.Add(1)

	go func() {
		min := 0
		var buffers []Entry

		for {
			select {
			case entry := <-outputWorker:
				if entry.LineNo > NullEntry.LineNo {
					buffers = Sort(append(buffers, entry))

					for i, entry := range buffers {
						if min+1 == entry.LineNo {
							fmt.Fprintln(iomap.Out, entry.Formatted)
							buffers = Nullify(buffers, i)
							min++
						}
					}

					buffers = Compact(buffers)
				} else {
					for i, entry := range buffers {
						fmt.Fprintln(iomap.Out, entry.Formatted)
						buffers = Nullify(buffers, i)
						min++
					}
					outputWait.Done()
					break
				}
			}
		}
	}()

	return &outputWait
}

func Sort(slice []Entry) []Entry {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].LineNo < slice[j].LineNo
	})

	return slice
}

func Nullify(slice []Entry, i int) []Entry {
	slice[i] = NullEntry
	return slice
}

func Compact(slice []Entry) []Entry {
	var new []Entry

	for _, entry := range slice {
		if entry.LineNo != NullEntry.LineNo {
			new = append(new, entry)
		}
	}

	return new
}
