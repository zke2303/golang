package main

import (
	"fmt"
	"jwt_demo/utils"
)

func main() {
	token, err := utils.GenerateJwt("caoyunkai")
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(token)

	claims, err := utils.ParseJwt(token)
	if err != nil {
		fmt.Println("parse fail")
		return
	}

	fmt.Println(claims.Username)
}
