package main

import "fmt"

func main() {
	defer fmt.Println("hello world")

	defer fmt.Println("hello Golang")

	defer fmt.Println("hello Python")

}
