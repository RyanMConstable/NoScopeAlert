package main

import (
	"fmt"
)

func CheckGameCodes(u Users) error {
	//Here we must check forever the status of the games of users.
	//Asynchronous will most likely be the best way to do this.
	for {
		fmt.Println("Never ending")
	}
	return nil
}
