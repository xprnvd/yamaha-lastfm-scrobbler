package main

import (
	"fmt"
)

func main() {

	get_yamaha()

	if auth_token == "" {
		fmt.Print("auth_token is empty")
		if session_key == "" {
			get_auth_token()
			fmt.Println(string(auth_token))
			get_authorization()
			// wait for input yes/no to continue
			var input string
			fmt.Scanln(&input)
			if input == "yes" {
				get_session()
			} else {
				fmt.Println("Exiting...")
			}

		}
	} else {
		get_sig()

	}
}
