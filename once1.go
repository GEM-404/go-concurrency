package main

import (
	"fmt"
	"sync"
)

var count int
var once sync.Once

func increment() {
	count++
}

func decrement() {
	count--
}

func main() {
	once.Do(increment)
	once.Do(decrement)

	fmt.Printf("Count: %d\n", count)
}
