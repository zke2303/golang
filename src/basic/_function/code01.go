package main

import "fmt"

func add[T int | float64](x, y T) T {
	return x + y
}

func main() {
	fmt.Println(add(1.2, 2.32))
}
