package main

import "fmt"

func main() {

	arr := make([]int, 0, 10)
	arr = append(arr, 1, 2, 3, 4)
	fmt.Println(len(arr))

	fmt.Println(arr[1:4])

	arr = append([]int{-2, -1, 0}, arr...)
	fmt.Println(arr)

}
