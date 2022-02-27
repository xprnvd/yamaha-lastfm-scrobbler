package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: yms [yamaha_IP] [lastfm_api_key] [lastfm_secret]")
		os.Exit(1)
	}
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
			if track_changed {
				fmt.Println("\nTrack changed, sending to lastfm...Next check in 60 Seconds")
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
