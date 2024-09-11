package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	allocationsCount := 0

	pool := sync.Pool{
		New: func() interface{} {
			fmt.Print(".")
			allocationsCount++
			return make([]byte, 1024)
		},
	}

	var waitGroup sync.WaitGroup

	for i := 0; i < 1000; i++ {
		waitGroup.Add(1)
		go func() {
			obj := pool.Get().([]byte)
			fmt.Print("-")
			// we add a bit of delay
			time.Sleep(100 * time.Millisecond)
			pool.Put(obj)
			waitGroup.Done()
		}()
		time.Sleep(10 * time.Millisecond)
	}

	waitGroup.Wait()

	fmt.Println("\n Number of allocations: ", allocationsCount)
}
