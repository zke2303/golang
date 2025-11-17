package main

import "fmt"

func main() {
	defer fmt.Println("Hello")
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	defer fmt.Println("Golang")
}
