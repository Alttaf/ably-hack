[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/alttaf/ably-hack.svg)](https://golang.org/)
[![Report](https://goreportcard.com/badge/github.com/alttaf/ably-hack?style=flat&v=1)](https://goreportcard.com/report/github.com/alttaf/ably-hack)

# Ably Hackathon Entry: Ably Text
```

 █████  ██████  ██      ██    ██     ████████ ███████ ██   ██ ████████ 
██   ██ ██   ██ ██       ██  ██         ██    ██       ██ ██     ██    
███████ ██████  ██        ████          ██    █████     ███      ██    
██   ██ ██   ██ ██         ██           ██    ██       ██ ██     ██    
██   ██ ██████  ███████    ██           ██    ███████ ██   ██    ██    
```

## About
This is an Ably GopherCon 2021 Hackathon entry based on [Teletext](https://en.wikipedia.org/wiki/Teletext).

Teletext was a protocol that allowed broadcasters to send text signals to televisions. This project is a similar concept, broadcasting tweet counts for a given hashtag in real time, via [Ably](https://ably.com)'s realtime API.

Visit the site at <https://mighty-lake-60203.herokuapp.com> to see the data in real time for the keyword "google" ([this search](https://twitter.com/search?q=google&src=typed_query&f=live) shows the equivalent data in Twitter). 
This keyword was chosen for the demo since there due to the reliability of tweets matching this keyword every minute! 

The table automatically updates: a new entry will be added at the top with each new minute, with data from the last two hours being shown.

This repo contains the backend code, see <https://github.com/Alttaf/ably-fe> for the frontend
![AblyTextScreenshot.png](AblyTextScreenshot.png)

### Inspiration

## Documentation Fixes

## Future work
This could further be expanded to:
- Have a search bar on the home page, so users could search for whichever hashtag they liked
- Track sentiment analysis of hashtags (events) in real-time:
  - Keep track of the sentiment of certain events
  - Monitor multiple events, see which has the highest overall positive sentiment, and most negative sentiment

## Running locally
1. Copy and paste `.env.example` file
2. Rename to `.env`
3. Fill in the specified values

## Contributors
1. [Alttaf Hussain](https://github.com/alttaf)
2. [Ed Harrod](https://github.com/echarrod)
