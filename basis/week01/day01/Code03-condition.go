package main

import "fmt"

func main() {

START:
	fmt.Println("start..")
	a := 10
	for a > 0 {
		fmt.Println("a = ", a)
		a--
	}
	goto START
	fmt.Println("Finish")
}
