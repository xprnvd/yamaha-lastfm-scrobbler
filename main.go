package main

import (
	"fmt"
	"time"
)

func main() {
	if auth_token == "" {
		if session_key == "" {
			get_auth_token()
			fmt.Println(string("\nNo Auth token store. New auth token: " + auth_token + "\n"))
			get_authorization()
			fmt.Println("\nOnce authorized type 'yes' to continue, or 'no' to exit")
			var input string
			fmt.Scanln(&input)
			if input == "yes" {
				get_session_sig()
				get_session()
			} else {
				fmt.Println("Exiting...")
			}
		}
	} else if session_key == "" {
		get_session_sig()
		get_session()
	} else {
		for {
			get_yamaha()
			track_check()
			if track_changed {
				get_scrobbler_sig()
				send_scrobbler()
			}
			time.Sleep(time.Minute * 1)
		}
	}
}
