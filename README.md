# Yamaha MusicCast LastFM DLNA Scrobber in GoLang

## What

Reads DLNA playback by making request to MusicCast server, and sends scrobbing data to LastFM account with minimal stdout.

## Why

I use plex for my music, and use it's last.fm integration for music discovery. For MusicCast, I use it's DLNA server, which does not have last.fm integration. So I needed to integrate the DLNA player with my last.fm account to get all play data.

## How

Current playing artist and track collection happens from MusicCast server. The application then connects to LastFM API via api.last.fm, and sends scrobble data.

### Required interaction

IP address of MusicCast server, LastFM API key pair generation and authorization.

Create api credentials: <https://www.last.fm/api/account/create>

Documentation on lastfm api: <https://www.last.fm/api/scrobbling>

## Build

After installing go, run the following command to build the application inside the repo directory:
```go build -o yls src/*.go```


API key 	4f5f0f0ae75b7b6c9c091dc6be31c1c5
Shared secret 	e1afe1329d5a054929e1eb3a1b8f6835