[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/alttaf/ably-hack.svg)](https://golang.org/)
[![Report](https://goreportcard.com/badge/github.com/alttaf/ably-hack?style=flat&v=1)](https://goreportcard.com/report/github.com/alttaf/ably-hack)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://github.com/alttaf/ably-hack/blob/main/LICENSE.md)

# Ably Hackathon Entry: Ably Text
```

 █████  ██████  ██      ██    ██     ████████ ███████ ██   ██ ████████ 
██   ██ ██   ██ ██       ██  ██         ██    ██       ██ ██     ██    
███████ ██████  ██        ████          ██    █████     ███      ██    
██   ██ ██   ██ ██         ██           ██    ██       ██ ██     ██    
██   ██ ██████  ███████    ██           ██    ███████ ██   ██    ██    
```

## About
This is an [Ably GopherCon 2021 Hackathon](https://github.com/ably-labs/Gophercon-Hackathon) entry based on [Teletext](https://en.wikipedia.org/wiki/Teletext).

Teletext was a protocol that allowed broadcasters to send text signals to televisions. This project is a similar concept, broadcasting tweet counts for a given hashtag in real time, via [Ably](https://ably.com)'s realtime API.

Visit the site at <https://mighty-lake-60203.herokuapp.com> to see the data in real time. You can enter a search term to see all the recent tweets matching that keyword.
The default display is for the keyword "google", due to this keyword reliably matching tweets every minute! ([this search](https://twitter.com/search?q=google&src=typed_query&f=live) shows the equivalent data in Twitter).

The table automatically updates in real time: a new entry will be added at the top as each new minute passes, with data from the last two hours being shown.

This repo contains the backend code, see <https://github.com/Alttaf/ably-fe> for the frontend
![AblyTextScreenshot.png](AblyTextScreenshot.png)

## Future work
This could further be expanded to:
- Track sentiment analysis of hashtags (events) in real-time:
  - Keep track of the sentiment of certain events
  - Monitor multiple events, see which has the highest overall positive sentiment, and most negative sentiment

## Running locally
1. Copy and paste `.env.example` file
2. Rename to `.env`
3. Fill in the specified values
4. Run `go run ably-text`

## Contributors
1. [Alttaf Hussain](https://github.com/alttaf)
2. [Ed Harrod](https://github.com/echarrod)
