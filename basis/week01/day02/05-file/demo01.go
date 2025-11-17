package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("D:\\user\\learning\\go\\workspace\\basis\\week01\\day02\\05-file\\text.txt", os.O_RDONLY, 0666)

	if err != nil {
		fmt.Println("打开文件错误")
		fmt.Println(err)
		return
	}
	buffer := make([]byte, 0, 512)

	offset, err := file.Read(buffer[:cap(buffer)])
	fmt.Println(offset)

	buffer = buffer[:offset]

	fmt.Println(string(buffer))

}
