package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var lastfm_api_key = "931461e741abc95aba1f6fe49fef306f"

// Response is the struct that will be unmarshalled from the JSON response
type Response struct {
	ResponseCode     int      `json:"response_code"`
	Input            string   `json:"input"`
	PlayQueueType    string   `json:"play_queue_type"`
	Playback         string   `json:"playback"`
	Repeat           string   `json:"repeat"`
	Shuffle          string   `json:"shuffle"`
	PlayTime         int      `json:"play_time"`
	TotalTime        int      `json:"total_time"`
	Artist           string   `json:"artist"`
	Album            string   `json:"album"`
	Track            string   `json:"track"`
	AlbumartURL      string   `json:"albumart_url"`
	AlbumartID       int      `json:"albumart_id"`
	UsbDevicetype    string   `json:"usb_devicetype"`
	AutoStopped      bool     `json:"auto_stopped"`
	Attribute        int      `json:"attribute"`
	RepeatAvailable  []string `json:"repeat_available"`
	ShuffleAvailable []string `json:"shuffle_available"`
}

type lastfm_token struct {
	Token string `json:"token"`
}

//function to handle status request from yamaha
func yamaha_request() {
	resp, err := http.Get("http://192.168.0.243:80/YamahaExtendedControl/v1/netusb/getPlayInfo")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	//fmt.Println(string(body))

	var result Response
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	fmt.Println(result.Artist, result.Album, result.Track, result.PlayTime)
}

// get auth token from lastfm
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

func main() {
	yamaha_request()
	p()
}
