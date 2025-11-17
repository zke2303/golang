package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("text.txt", os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("打开文件错误")
		return
	}

	offset, err := file.WriteString("Hello Golang")
	if err != nil {
		fmt.Println("写入文件错误")
		return
	}
	fmt.Println(offset)
	file.Close()
}
