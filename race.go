package main

import (
	"fmt"
	"sync"
	"time"
)

var data int

func main0() {
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("The value is %v.\n", data)
	}
}

func main1() {
	go func() { data++ }()
	time.Sleep(1 * time.Second)
	if data == 0 {
		fmt.Printf("The value is %v.\n", data)
	}
}

func main2() {
	go func() { data++ }()
	if data == 0 {
		fmt.Println("The value is 0.")
	} else {
		fmt.Printf("the value is %v.\n", data)
	}
}

func main3() {
	var memoryAccess sync.Mutex
	var value int
	go func() {
		memoryAccess.Lock()
		value++
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if value == 0 {
		fmt.Printf("the value is %v.\n", value)
	} else {
		fmt.Printf("the value is %v.\n", value)
	}
	memoryAccess.Unlock()
}
