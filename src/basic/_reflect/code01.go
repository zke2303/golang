package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "hello golang"
	reflectType := reflect.TypeOf(str)
	fmt.Println(reflectType)
	fmt.Println(reflectType.PkgPath())
}
