package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

var key_after_hash = ""

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func get_sig() {
	fmt.Println(string("Api Key: " + api_key))
	fmt.Println(string("Api Sig: " + api_sig))
	fmt.Println(string("Auth Token: " + auth_token))
	var key_before_hash = "api_key" + api_key + "methodauth.getSessiontoken" + auth_token
	fmt.Println(string("Key Before Hash: " + key_before_hash))
	var md5_hash = GetMD5Hash(key_before_hash)
	fmt.Println(string("Api Sig: " + md5_hash))
	key_after_hash = "api_key" + api_key + "methodauth.getSessiontoken" + auth_token + "api_sig" + md5_hash + "&format=json"
	fmt.Println(string("Key After Hash: " + key_after_hash))
}
