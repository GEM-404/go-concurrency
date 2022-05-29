package main

import "sync"

type counter struct {
	mu    sync.Mutex
	value int
}

func (c *counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}
