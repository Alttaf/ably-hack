package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ably/ably-go/ably"
	"github.com/julienschmidt/httprouter"
)

type TwitterData struct {
	Data []struct {
		End        time.Time `json:"end"`
		Start      time.Time `json:"start"`
		TweetCount int       `json:"tweet_count"`
	} `json:"data"`
	Meta struct {
		TotalTweetCount int `json:"total_tweet_count"`
	} `json:"meta"`
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	ablyClient, err := ably.NewRealtime(ably.WithKey("J0HbFg.02ttpg:pJoo1J-jZV2Hym28kvuNKQwopg66c9bB9SIXDUhAMFw"))
	if err != nil {
		fmt.Println("Dying could not create client")
		return
	}
	twitterQuery := ps.ByName("name")
	channel := ablyClient.Channels.Get("test")

	/* Publish a message to the test channel */
	ctx, _ := context.WithTimeout(context.Background(), 120*time.Second)

	// defer canceling so that all the resources are freed up
	// For this and the derived contexts
	defer func() {
		//fmt.Println("Main Defer: canceling context")
		//cancelFunction()
	}()

	go func() {
		for i := 0; i < 5; i++ {
			data, err := callTwitter(twitterQuery)
			if err != nil {
				fmt.Println("could not get twitter data")
			}
			fmt.Println(data)
			err = channel.Publish(ctx, "test", data)
			if err != nil {
				fmt.Printf("there was an err %v", err)
			}
			time.Sleep(20 * time.Second)
		}
	}()
	// w.Header().Add()

	fmt.Fprintf(w, "Hello Ably, %s!\n", ps.ByName("name"))
}

func callTwitter(query string) (string, error) {
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

	req.Header.Add("Authorization", "Bearer AAAAAAAAAAAAAAAAAAAAABKhVAEAAAAAIDysETWcnW%2FP6b8ufmCigkDiKc0%3Dgd0OrHm53uaMPrJy289NtjHCVEUVDOMsORX92dgRRjr4TlB7Dd")
	req.Header.Add("Cookie", "personalization_id=\"v1_R/HwN1zpTepqYV7B7JG1Ng==\"; guest_id=v1%3A163528855620592770")

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
	var tw TwitterData
	err = json.Unmarshal(body, &tw)
	if err != nil {
		log.Fatal("could not unmarshall")
	}
	fmt.Println("getting data from twitter")
	//fmt.Printf("data : %v", tw.Data)

	return string(body), nil
}

func main() {
	//os.Setenv("PORT", "8080")
	fmt.Printf("starting Server on port %s", ":"+os.Getenv("PORT"))

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
