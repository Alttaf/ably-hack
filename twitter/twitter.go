package twitter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"ably-text/env"
)

type Response struct {
	Data []struct {
		End        time.Time `json:"end"`
		Start      time.Time `json:"start"`
		TweetCount int       `json:"tweet_count"`
	} `json:"data"`
	Meta struct {
		TotalTweetCount int `json:"total_tweet_count"`
	} `json:"meta"`
}

func RetrieveTweets(query string) (string, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	now := time.Now()
	twoHours := time.Hour * time.Duration(-2)
	twoHoursAgo := now.Add(twoHours)
	t := twoHoursAgo.Format("2006-01-02T15:04:05Z")
	req, err := http.NewRequest("GET", "https://api.twitter.com/2/tweets/counts/recent?query="+query+"&granularity=minute&start_time="+t, nil)
	if err != nil {
		_ = fmt.Errorf("got error %s", err.Error())
		fmt.Println("Dying on making request")
		return "", nil
	}

	twitBearerToken := env.RetrieveValue("TWITTER_BEARER_TOKEN")
	if twitBearerToken == "" {
		return "", nil
	}

	twitPersonalisationId := env.RetrieveValue("TWITTER_PERSONALISATION_ID")
	if twitPersonalisationId == "" {
		return "", nil
	}
	twitGuestId := env.RetrieveValue("TWITTER_GUEST_ID")
	if twitGuestId == "" {
		return "", nil
	}

	req.Header.Add("Authorization", "Bearer "+twitBearerToken)

	cookieHeaderValue := fmt.Sprintf("personalization_id=\"%v\"; guest_id=%v", twitPersonalisationId, twitGuestId)
	req.Header.Add("Cookie", cookieHeaderValue)

	res, err := client.Do(req)
	if err != nil {
		_ = fmt.Errorf("got error %s", err.Error())
		fmt.Println("Dying on request")
		return "", nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	var tr Response
	err = json.Unmarshal(body, &tr)
	if err != nil {
		log.Fatal("could not unmarshall")
	}
	fmt.Println("getting data from twitter")

	return string(body), nil
}
