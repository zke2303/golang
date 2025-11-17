package main

import "fmt"

type User interface {
	GetId() uint64
	SetId(uint64)

	GetUsername() string
	SetUsername(string)

	GetPassword() string
	SetPassword(string)

	GetGender() string
	SetGender(string)

	GetAge() uint8
	SetAge(uint8)
}

type Role struct {
	id       uint64
	username string
	password string
	gender   string
	age      uint8
}

func (r *Role) GetId() uint64 {
	return r.id
}

func (r *Role) SetId(id uint64) {
	r.id = id
}

func main() {
	role := Role{
		id:       0001,
		username: "caoyunkai",
		password: "123456",
		gender:   "male",
		age:      21,
	}

	role.SetId(0002)

	fmt.Println(role)
}
