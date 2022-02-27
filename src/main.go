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
		fmt.Println("\nConstant checks every 60 Seconds. No further output except for errors. Review activity on last.fm console")
		for {
			get_yamaha()
			track_check()
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
