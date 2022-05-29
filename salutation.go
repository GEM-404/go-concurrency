package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	salutation := "Hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "Welcome"
	}()
	wg.Wait()
	fmt.Println(salutation)
}
