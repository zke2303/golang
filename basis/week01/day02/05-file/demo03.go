package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("text.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("打开文件异常")
		return
	}
	// 创建一个byte切片,用来存储数据
	buffer := make([]byte, 512, 512)
	offset, err := file.ReadAt(buffer, 2)
	fmt.Println(offset)
	fmt.Println(string(buffer))

}
