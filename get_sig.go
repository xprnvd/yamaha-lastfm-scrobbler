package main

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

var api_sig = ""
var sk_sig = ""

func get_session_sig() {
	var hash = "api_key" + api_key + "methodauth.getSessiontoken" + auth_token + api_sec
	api_sig = GetMD5Hash(hash)
}

func get_scrobbler_sig() {
	var nowtime = strconv.Itoa(int(time.Now().Unix()))
	var hash = "api_key" + api_key + "artist" + artist + "methodtrack.scrobblesk" + session_key + "timestamp" + nowtime + "track" + track + api_sec
	sk_sig = GetMD5Hash(hash)
}
