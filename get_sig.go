package main

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

var api_sig = ""

func get_sig() {
	//fmt.Println(string("Api Key: " + api_key))
	//fmt.Println(string("Api Sig: " + api_sec))
	//fmt.Println(string("Auth Token: " + auth_token))
	var hash = "api_key" + api_key + "methodauth.getSessiontoken" + auth_token + api_sec
	//fmt.Println(string("Key Before Hash: " + hash))
	api_sig = GetMD5Hash(hash)
	//fmt.Println(string("Api Sig: " + api_sig))

}
