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

## Build & Run

After installing go, download the release or run the following command to build from source:
```go build -o yls src/*.go```

Run with tmux or zellij for a detached session

## Output

```txt
[host ~]# yamaha-lastfm-scrobbler/bin/yls 
IP Address: 192.168.0.243
API Key: xxxx
API Secret: xxxx

Please go to http://www.last.fm/api/auth/?api_key=xxxx&token=xxxx to authorize this app

Once authorized type 'yes' to continue, or 'no' to exit

yes

Constant checks every 60 Seconds. No further output except for errors. Review activity on last.fm console
```
