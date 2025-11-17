package main

import "fmt"

type Author struct {
	name   string
	gender int8
	age    int8
}
type Book struct {
	title       string
	author      Author
	description string
	price       float64
}

func main() {
	author := Author{
		"江南",
		1,
		44,
	}

	book := Book{
		"龙族",
		author,
		"热血",
		41.99,
	}

	fmt.Println(book.author)
}
