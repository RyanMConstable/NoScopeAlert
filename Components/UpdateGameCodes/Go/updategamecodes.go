package main

import (
	"fmt"
)

func main() {
	var users Users

	InitUsers(&users)

	fmt.Println(users)
}

type Users struct {
	users []User
}

type User struct {
	key        string
	steamid    string
	steamidkey string
	knowncode  string
}
