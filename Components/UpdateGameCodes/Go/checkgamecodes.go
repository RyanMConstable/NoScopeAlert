package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Response struct {
	Result struct {
		Nextcode string `json:"nextcode"`
	} `json:"result"`
}

func CheckGameCodes(u *Users) error {
	//Here we must check forever the status of the games of users.
	//Asynchronous will most likely be the best way to do this.
	for {
		//Loop through users
		for i, user := range u.users {
			gameCodeURL := fmt.Sprintf("https://api.steampowered.com/ICSGOPlayers_730/GetNextMatchSharingCode/v1?key=%v&steamid=%v&steamidkey=%v&knowncode=%v", user.key, user.steamid, user.steamidkey, user.knowncode)

			resp, err := http.Get(gameCodeURL)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			var data Response
			err = json.Unmarshal(body, &data)
			if err != nil {
				return err
			}

			//Here we need to send the nextcode IF it does not equal n/a to rabbitmq queue
			if data.Result.Nextcode != "n/a" {
				fmt.Println("A new code has been found: ", data.Result.Nextcode)

				//This sends the info to the rabbitmq queue to download the game
				//NOTE: We do not check to see if this exists already, the workers will check that.
				SendToQueue(data.Result.Nextcode, user.steamid)

				//TMP We need to update the game code in the struct/db so that it keeps updating to new codes we cannot have stale data
				u.users[i].knowncode = data.Result.Nextcode
			}
		}
		time.Sleep(2 * time.Second)
	}
}
