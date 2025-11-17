package main

import "fmt"

func Max[T int | float64](arr []T) T {
	var maxValue T = arr[0]
	for _, value := range arr {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func main() {
	arr := []int{1, 2, 4, 5, 7, 1, 2, 20}
	fmt.Println(Max(arr))
}
