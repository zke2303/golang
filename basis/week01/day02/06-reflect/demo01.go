package main

import (
	"fmt"
	"reflect"
)

type Phone struct {
	username string
}

type myInt int

func main() {

	myMap := map[string]int{
		"username": 123,
	}
	rType := reflect.TypeOf(myMap)
	fmt.Println(rType.Elem().Size())
}
