package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"ably-text/env"
	"ably-text/twitter"

	"github.com/ably/ably-go/ably"
	_ "github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	ablyAPIKey := env.RetrieveValue("ABLY_API_KEY")
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
			data, err := twitter.RetrieveTweets(twitterQuery)
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
