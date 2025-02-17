package main

import (
	"fmt"
	"runtime/debug"
)

func weirdFunction() {
	defer recoverPanic()
	var numSl []int
	numSl = append(numSl, 1)
	numSl = append(numSl, 2)

	fmt.Println(numSl[100])
	fmt.Println("정상실행")
}

func recoverPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in f", r)
		debug.PrintStack()
	}
}

func main() {
	weirdFunction()
	fmt.Println("패닉후 함수...")
}
