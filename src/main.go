package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	get_values()
	get_auth_token()
	get_authorization()
	fmt.Println("\nOnce authorized type 'yes' to continue, or 'no' to exit")
	var input string
	fmt.Scanln(&input)
	if input == "yes" {
		get_session_sig()
		get_session()
		for {
			get_yamaha()
			track_check()
			fmt.Println("\nConstant checks every 60 Seconds")
			if track_changed {
				get_scrobbler_sig()
				send_scrobbler()
			}
			time.Sleep(time.Minute * 1)
		}
	} else {
		fmt.Println("Exiting...")
		os.Exit(0)
	}
}
