package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("text.txt")
	if err != nil {
		return
	}
	fmt.Println(string(data))
}
