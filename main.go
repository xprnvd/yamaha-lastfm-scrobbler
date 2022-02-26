package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

func yamaha_request() {
	resp, err := http.Get("http://192.168.0.243:80/YamahaExtendedControl/v1/netusb/getPlayInfo")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	fmt.Println(string(body))

	var result Response
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	fmt.Println(result.Artist)
	fmt.Println(result.Album)
	fmt.Println(result.Track)
	fmt.Println(result.PlayTime)
}

func main() {
	yamaha_request()
}
