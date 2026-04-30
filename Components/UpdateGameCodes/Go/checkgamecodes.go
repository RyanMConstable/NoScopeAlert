package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func CheckGameCodes(u Users) error {
	//Here we must check forever the status of the games of users.
	//Asynchronous will most likely be the best way to do this.
	for {
		//Loop through users
		for _, user := range u.users {
			gameCodeURL := fmt.Sprintf("https://api.steampowered.com/ICSGOPlayers_730/GetNextMatchSharingCode/v1?key=%v&steamid=%v&steamidkey=%v&knowncode=%v", user.key, user.steamid, user.steamidkey, user.knowncode)

			resp, err := http.Get(gameCodeURL)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			fmt.Println(string(body))

		}
		time.Sleep(2 * time.Second)
	}
}
