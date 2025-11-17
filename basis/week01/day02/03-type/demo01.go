package main

import (
	"fmt"
)

type userIdInt uint32
type orderInInt uint32

func main() {

	a := 10.10
	var b interface{} = a

	value, ok := b.(float64)
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("type of a is not int")
	}

}
