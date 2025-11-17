package main

import "fmt"

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func sum(begin, end int) int {
	sum := 0
	for begin <= end {
		sum += begin
		begin++
	}
	return sum
}

func main() {
	a, b := 10, 9
	fmt.Println(max(a, b))

	fmt.Println(sum(1, 100))

	b, a = a, b
	fmt.Println(a, b)
}
