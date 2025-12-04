package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	wait.Add(1)
	go func() {
		fmt.Println(1)
		wait.Done()
	}()

	wait.Wait()
	fmt.Println(2)
}
