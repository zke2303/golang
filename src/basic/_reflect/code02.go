package main

import (
	"fmt"
	"reflect"
)

func main(){
	str := "hello world"
	value := reflect.ValueOf(str)
	fmt.Printf("type: %T\n", value)
}
