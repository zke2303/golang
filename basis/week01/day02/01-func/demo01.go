package main

import "fmt"

func Sum(args ...int) int {
	var sum int
	for _, value := range args {
		sum += value
	}
	return sum
}

func main() {
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
