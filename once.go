package main

import (
	"fmt"
	"sync"
)

var count int
var once sync.Once
var increments sync.WaitGroup

func increment() {
	count++
}

func main() {
	increments.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}

	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}
