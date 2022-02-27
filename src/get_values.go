package main

import "fmt"

func prompt(s string) string {
	fmt.Printf("%s: ", s)
	var input string
	fmt.Scanln(&input)
	return input
}

func get_values() {
	ip = prompt("IP Address")
	api_key = prompt("API Key")
	api_sec = prompt("API Secret")

}
