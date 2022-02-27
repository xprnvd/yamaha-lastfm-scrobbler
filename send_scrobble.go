package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func send_scrobbler() {
	var nowtime = strconv.Itoa(int(time.Now().Unix()))
	endpoint := "http://ws.audioscrobbler.com/2.0/"
	data := url.Values{}
	data.Set("api_key", api_key)
	data.Set("artist", artist)
	data.Set("method", "track.scrobble")
	data.Set("sk", session_key)
	data.Set("timestamp", nowtime)
	data.Set("track", track)
	data.Set("api_sig", sk_sig)

	client := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Status)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}
