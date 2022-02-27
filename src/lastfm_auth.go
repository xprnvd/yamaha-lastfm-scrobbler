package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var api_key = os.Args[2]
var api_sec = os.Args[3]
var auth_token = ""
var session_key = ""

type lastfm_token struct {
	Token string `json:"token"`
}

type lastfm_session struct {
	Session struct {
		Name       string `json:"name"`
		Key        string `json:"key"`
		Subscriber int    `json:"subscriber"`
	} `json:"session"`
}

func get_auth_token() {
	resp, err := http.Get("http://ws.audioscrobbler.com/2.0/?method=auth.gettoken&api_key=" + api_key + "&format=json")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		fmt.Println("Can not read response body")
	}

	var result lastfm_token
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	auth_token = result.Token
}

func get_authorization() {
	fmt.Println("\nPlease go to http://www.last.fm/api/auth/?api_key=" + api_key + "&token=" + auth_token + " to authorize this app")
}

func get_session() {
	resp, err := http.Get("http://ws.audioscrobbler.com/2.0/?method=auth.getSession&api_key=" + api_key + "&token=" + auth_token + "&format=json&api_sig=" + api_sig)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Can not read response body")
	}

	var result lastfm_session
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if strings.Contains(string(body), "error") {
		fmt.Println("Session error. Check for authorization upon retry")
		os.Exit(1)
	}
	session_key = result.Session.Key
}
