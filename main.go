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
	_ "github.com/joho/godotenv"
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
	ablyAPIKey := retrieveEnvValue("ABLY_API_KEY")
	if ablyAPIKey == "" {
		return
	}

	ablyClient, err := ably.NewRealtime(ably.WithKey(ablyAPIKey))
	if err != nil {
		fmt.Println("error when attempting to create client :: " + err.Error())
		return
	}
	twitterQuery := ps.ByName("name")
	channel := ablyClient.Channels.Get("test")

	/* Publish a message to the test channel */
	ctx, _ := context.WithTimeout(context.Background(), 120*time.Second)
	// TODO: the cancel function returned by context.WithTimeout should be called, not discarded, to avoid a context leak

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
	fmt.Fprintf(w, "Hello Ably, %s!\n", ps.ByName("name"))
}

func retrieveEnvValue(envKey string) string {
	envValue, present := os.LookupEnv(envKey)
	if !present {
		errorString := fmt.Errorf("%v unset. This must be set from an .env file", envKey)
		fmt.Println(errorString)
		return ""
	}

	return envValue
}

// TODO: move to subpackage
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

	twitBearerToken := retrieveEnvValue("TWITTER_BEARER_TOKEN")
	if twitBearerToken == "" {
		return "", nil
	}

	twitPersonalisationId := retrieveEnvValue("TWITTER_PERSONALISATION_ID")
	if twitPersonalisationId == "" {
		return "", nil
	}
	twitGuestId := retrieveEnvValue("TWITTER_GUEST_ID")
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
	var tw TwitterData
	err = json.Unmarshal(body, &tw)
	if err != nil {
		log.Fatal("could not unmarshall")
	}
	fmt.Println("getting data from twitter")

	return string(body), nil
}

func main() {
	port, present := os.LookupEnv("PORT")
	if !present {
		port = "8080"
	}

	fmt.Printf("starting Server on port %s", ":"+port)

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
