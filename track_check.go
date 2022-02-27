package main

var last_track = ""
var track_changed = false

func track_check() {
	if last_track != track {
		track_changed = true
		last_track = track
	} else {
		track_changed = false
	}
}
