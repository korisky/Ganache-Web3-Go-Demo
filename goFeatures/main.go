package main

import (
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	count := 1_000_000

	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Second)
		}()
	}

	wg.Wait()
}
