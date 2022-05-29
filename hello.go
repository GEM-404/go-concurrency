package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	start := time.Now()

	wg.Add(3)

	go sayHello()
	// continue doing other things

	go func() {
		defer wg.Done()
		fmt.Println("Hello")
		time.Sleep(3 * time.Second)
		fmt.Println("Yaaawn, I'm waking up from a 'hello' sleep")
	}()

	sayBye := func() {
		defer wg.Done()
		fmt.Println("Goodbye")
		time.Sleep(3 * time.Second)
		fmt.Println("Yaaawn, I'm waking up from a 'Goodbye' sleep")
	}

	go sayBye()

	wg.Wait()

	end := time.Since(start)

	fmt.Printf("Took %v\n", end)
}

func sayHello() {
	defer wg.Done()
	fmt.Println("hi")
	time.Sleep(3 * time.Second)
	fmt.Println("Yaaawn, I'm waking up from a 'hi' sleep")
}
