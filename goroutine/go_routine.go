package main

import (
	"fmt"
	"sync"
)

func run(wg *sync.WaitGroup, startNumber int) {
	fmt.Println("고루틴 추가")
	for i := startNumber; i < startNumber+100; i++ {
		fmt.Println(i)
	}
	wg.Done()
	fmt.Println("고루틴 종료")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go run(&wg, -100)
	go run(&wg, 0)
	go run(&wg, 100)
	wg.Wait()
}
