package main

import "fmt"

func main(){
	var intCh chan int
	go func(){
		intCh <- 10
	}()
	num := <- intCh
	fmt.Println(num)

}

