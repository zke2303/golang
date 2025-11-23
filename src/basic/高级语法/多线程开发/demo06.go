package main
import (
	"fmt"
	"sync"
)

func main(){
	var wait sync.WaitGroup
	wait.Add(3)
	go func(){
		fmt.Println(1)
		wait.Done()
	}()
	go func(){
		fmt.Println(2)
		wait.Done()
	}()
	go func(){
		fmt.Println(3)
		wait.Done()
	}()

	wait.Wait()
	fmt.Println("main")
}

