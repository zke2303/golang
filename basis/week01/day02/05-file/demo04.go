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

	buffer := make([]byte, 0, 512)

	offset, err := file.Read(buffer[len(buffer):cap(buffer)])
	fmt.Println(string(buffer))
	fmt.Println(string(buffer[:offset]))
	buffer = buffer[:len(buffer)+offset]
	fmt.Println(string(buffer))

}
