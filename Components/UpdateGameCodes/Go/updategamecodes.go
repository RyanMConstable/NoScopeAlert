package main

import (
	"fmt"
	"log"
)

func main() {
	var users Users

	//Function to set users
	//In the future this will be a database grab, and then the loop will sit in memory for ever.
	//Adding new users must be instant and will have to be thought about as well
	//For now this is just my user for testing (not included in git so people don't steal my keys)
	fmt.Println("Initializing users")
	InitUsers(&users)

	//Here we must constantly check every user to see if a new game has been played.
	//If a game has been played, then we must send that code and the steamid to a rabbitmq queue
	fmt.Println("Checking players for new games played")
	err := CheckGameCodes(users)
	if err != nil {
		log.Fatal(err)
	}

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
