package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() { sayhello() }

func sayhello() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("hello")
	}()

	wg.Wait()
}

