package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var lastfm_api_key = "931461e741abc95aba1f6fe49fef306f"

type lastfm_token struct {
	Token string `json:"token"`
}

func lastfm_auth() {
	resp, err := http.Get("http://ws.audioscrobbler.com/2.0/?method=auth.gettoken&api_key=" + lastfm_api_key + "&format=json")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	fmt.Println(string(body))

	var result lastfm_token
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	fmt.Println(result.Token)

}
