// APIs we are planning to call
//
// 1. /top-headlines/in
// 2. /top-headlines/at
// 3. /get-everything/bangalore
// 4. /get-everything/innsbruck
// handleTopHeadlines handles top-headlines route
// handleGetEverything handles get-everything route
//
// ..

package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/richarddes/newsapi-golang"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print(w, "Hello, welcome to news service")
	})

	router.HandleFunc("/top-headlines/{country}", handleTopHeadlines)
	router.HandleFunc("/get-everything/{q}", handleGetEverything)

	log.Println("Listening on port 8000")
	http.ListenAndServe(":8000", router)
}

func handleTopHeadlines(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming top headlines request")
	ctx, vars := ctxVars(r)

	opts := newsapi.TopHeadlinesOpts{
		Country: vars["country"],
	}

	c := newsApiClient()
	resp, err := c.TopHeadlines(ctx, opts)

	if err != nil {
		log.Fatal(err)
	}

	printResponse(w, resp)
}

func handleGetEverything(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming get everything request")
	ctx, vars := ctxVars(r)

	opts := newsapi.EverythingOpts{
		Q:      vars["q"],
		SortBy: "popularity",
	}

	c := newsApiClient()

	resp, err := c.Everything(ctx, opts)

	if err != nil {
		log.Fatal(err)
	}

	res := newsapi.TopHeadlinesResp(resp)

	printResponse(w, res)
}

func ctxVars(r *http.Request) (context.Context, map[string]string) {
	ctx := context.Background()
	vars := mux.Vars(r)
	return ctx, vars
}

func printResponse(w http.ResponseWriter, resp newsapi.TopHeadlinesResp) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func newsApiClient() newsapi.Client {
	return newsapi.Client{APIKey: os.Getenv("NEWS_API_KEY")}
}
