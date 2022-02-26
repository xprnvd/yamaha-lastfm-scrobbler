package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

//function to handle status request from yamaha
func get_yamaha() {
	resp, err := http.Get("http://192.168.0.243:80/YamahaExtendedControl/v1/netusb/getPlayInfo")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		fmt.Println("Can not read response body")
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	//fmt.Println(string(body))
	fmt.Println("\nCurrent playback from Yamaha: ", result.Artist, result.Album, result.Track, result.PlayTime)
}
