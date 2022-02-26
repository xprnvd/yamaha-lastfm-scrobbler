package main

import (
	"fmt"
)

func main() {

	get_yamaha()

	if auth_token == "" {
		if session_key == "" {
			get_auth_token()
			fmt.Println(string("\nNo Auth token store. New auth token: " + auth_token + "\n"))
			get_authorization()
			fmt.Println("\nOnce authorized type 'yes' to continue, or 'no' to exit")
			var input string
			fmt.Scanln(&input)
			if input == "yes" {
				get_sig()
				get_session()
			} else {
				fmt.Println("Exiting...")
			}

		}
	} else {
		get_sig()
		get_session()

	}
}
