package io

import (
	"bufio"
	"fmt"
	"runtime"
	"sync"
)

func Worker(iomap IOMap, formatter func(Entry) (Entry, bool)) {
	outputWorker := make(chan Entry)
	outputWait := output(&iomap, outputWorker)
	reader(&iomap, outputWorker, formatter)
	outputWait.Wait()
}

func reader(iomap *IOMap, outputWorker chan Entry, formatter func(Entry) (Entry, bool)) {
	var readWait sync.WaitGroup

	scanner := bufio.NewScanner(iomap.In)
	count := 0
	limit := runtime.NumCPU()
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
			entry, ok := formatter(entry)

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

func output(iomap *IOMap, outputWorker chan Entry) *sync.WaitGroup {
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
