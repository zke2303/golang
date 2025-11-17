package main

import (
	"fmt"
	"reflect"
)

func main() {
	num := 1000
	ele := reflect.ValueOf(&num).Elem()
	fmt.Println(ele)
}
