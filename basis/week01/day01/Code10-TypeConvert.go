package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 10
	fmt.Println(float64(a))
	fmt.Printf("typeOf(%T)\n", float64(a))

	var s string = "10"
	fmt.Println(strconv.Atoi(s))

	fmt.Println(strconv.Itoa(a))

	var b rune = 'a'
	fmt.Println(b)

}
