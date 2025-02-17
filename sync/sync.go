package main

import (
	"fmt"
	"sync"
)

var rwMu sync.RWMutex
var data int

func read() {
	rwMu.RLock()
	defer rwMu.RUnlock()
	fmt.Println("읽기:", data)
}

func write(newValue int) {
	rwMu.Lock()
	defer rwMu.Unlock()
	data = newValue
	fmt.Println("쓰기:", data)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()
		read()
	}()
	go func() {
		defer wg.Done()
		write(1234)
	}()
	go func() {
		defer wg.Done()
		read()
	}()

	wg.Wait()
}
