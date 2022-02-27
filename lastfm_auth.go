package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var api_key = "931461e741abc95aba1f6fe49fef306f"
var api_sec = "b68203de996c0e0cad617ce124b64a20"
var auth_token = "DNDCQ4Uva7kdPlk9uMeceYpOXFX-q_SY"
var session_key = "mrYNC7UqH0nzgveyfDXFTQn-7AxgfjUl"

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
	fmt.Println("Please go to http://www.last.fm/api/auth/?api_key=" + api_key + "&token=" + auth_token + " to authorize this app")
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
	session_key = result.Session.Key
	fmt.Println(string(session_key))
}
