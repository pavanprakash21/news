package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	tr "github.com/pavanprakash21/news/pkg/translate"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print(w, "Hello, welcome to translation service")
	})

	router.HandleFunc("/translate", handleTranslation)

	http.ListenAndServe(":8001", router)
}

func handleTranslation(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	type Input struct {
		Text *string `json:"text"`
	}

	var input Input

	err := decoder.Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if input.Text == nil {
		http.Error(w, "missing field 'text' from JSON object", http.StatusBadRequest)
		return
	}

	if decoder.More() {
		http.Error(w, "extraneous data after JSON object", http.StatusBadRequest)
		return
	}

	translate := tr.New(tr.Config{
		Url: os.Getenv("TRANSLATION_SERVICE_URL"),
		Key: os.Getenv("TRANSLATION_SERVICE_KEY"),
	})

	resp, err := translate.Translate(*input.Text, "de", "en")

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
