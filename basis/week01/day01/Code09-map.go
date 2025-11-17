package main

import "fmt"

func main() {
	myMap := map[string]string{
		"username": "曹云开",
		"gender":   "male",
		"age":      "21",
	}

	for key, value := range myMap {
		fmt.Println("key = ", key, ", value = ", value)
	}

	fmt.Println(myMap["username"])

	fmt.Println(len(myMap))

	delete(myMap, "gender")

	fmt.Println(myMap)
}
