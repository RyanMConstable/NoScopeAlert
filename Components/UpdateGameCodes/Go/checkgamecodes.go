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
			if err != nil {
				return err
			}

			var data Response
			err = json.Unmarshal(body, &data)
			if err != nil {
				return err
			}

			fmt.Println(data.Result.Nextcode)

			//Here we need to send the nextcode IF it does not equal n/a to rabbitmq queue
		}
		time.Sleep(2 * time.Second)
	}
}
