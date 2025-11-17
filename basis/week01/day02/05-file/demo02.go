package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func ReadFile(file *os.File) ([]byte, error) {
	// 创建一个byte数组,用来存在数据
	buffer := make([]byte, 0, 512) // 初始化容量为512, 数据长度为0
	for {
		// 如果buffer容量不足是, 进行扩容
		if len(buffer) == cap(buffer) {
			buffer = append(buffer, 0)[:len(buffer)]
		}
		// 继续读取文件内容
		offser, err := file.Read(buffer[len(buffer):cap(buffer)])
		// 将已经写入的数据归入切片
		buffer = buffer[:len(buffer)+offser]
		// 发生错误是
		if err != nil {
			if errors.Is(err, io.EOF) {
				err = nil
			}
			return buffer, err
		}

	}
}

func main() {
	data, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println("读取文件失败", err)
		return
	}
	fmt.Println(string(data))
}