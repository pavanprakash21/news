package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"log"
	"net/http"
	"sort"

	concurrently "github.com/pavanprakash21/news/pkg/concurrently"
	"github.com/pavanprakash21/news/pkg/types"
)

var NewsLinks = []string{
	"https://newsapi.org/v2/top-headlines?country=in&apiKey=" + os.Getenv("NEWS_API_KEY"),
	"https://newsapi.org/v2/top-headlines?country=at&apiKey=" + os.Getenv("NEWS_API_KEY"),
	"https://newsapi.org/v2/everything?q=bangalore&apiKey=" + os.Getenv("NEWS_API_KEY"),
	"https://newsapi.org/v2/everything?q=innsbruck&apiKey=" + os.Getenv("NEWS_API_KEY"),
}

var NewsTopics = []string{"India", "Austria", "Bangalore", "Innsbruck"}

type data struct {
	News []types.NewsResult `json:"news"`
}
type finalData struct {
	Data data `json:"data"`
}

func main() {
	res := concurrently.MakeRequests(NewsLinks, 4)

	for idx := range res {
		if idx == 1 || idx == 3 {
			articles := res[idx].Articles
			res[idx].Articles = translated_articles(articles, 8)
		}
		res[idx].Topic = NewsTopics[idx]
	}

	data := data{res}
	finalResp := finalData{data}

	body, err := json.MarshalIndent(finalResp, "", "  ")
	panicIf(err)

	filename := time.Now().Format("2006-01-02") + ".json"
	filename = "data/" + filename
	err = ioutil.WriteFile(filename, body, 0644)
	panicIf(err)
	log.Println("Wrote to the file")
}

func translated_articles(articles []types.Article, concurrencyLimit uint8) []types.Article {
	type articleMap struct {
		index   int
		article types.Article
	}

	semaphoreChan := make(chan struct{}, concurrencyLimit)
	resultsChan := make(chan *articleMap)

	defer func() {
		close(semaphoreChan)
		close(resultsChan)
	}()

	for idx, article := range articles {
		go func(idx int, article types.Article) {
			semaphoreChan <- struct{}{}
			translated := translatedJSON(&article)
			raw := []byte(translated)
			json.Unmarshal(raw, &article)
			resp := &articleMap{idx, article}
			resultsChan <- resp
			<-semaphoreChan
		}(idx, article)
	}

	var resps []articleMap

	for {
		resp := <-resultsChan
		resps = append(resps, *resp)

		if len(resps) == len(articles) {
			break
		}
	}

	sort.Slice(resps, func(i, j int) bool {
		return resps[i].index < resps[j].index
	})

	var results []types.Article

	for _, resp := range resps {
		results = append(results, resp.article)
	}

	return results
}

func toJSON(value interface{}) string {
	bytes, err := json.Marshal(value)
	panicIf(err)
	return string(bytes)
}

func translatedJSON(article *types.Article) string {
	type Alias types.Article

	return toJSON(&struct {
		*Alias
		Title       string `json:"title"`
		Description string `json:"description"`
		Content     string `json:"content"`
	}{
		Alias:       (*Alias)(article),
		Title:       translateText(article.Title),
		Description: translateText(article.Description),
		Content:     translateText(article.Content),
	})
}

func translateText(text string) string {
	if text == "" {
		return ""
	}

	const (
		SourceLang = "de"
		TargetLang = "en"
	)

	reqBody, err := json.Marshal(map[string]string{
		"q":      text,
		"source": SourceLang,
		"target": TargetLang,
	})
	panicIf(err)

	byteBody := bytes.NewBuffer(reqBody)
	res, err := http.Post("http://localhost:5000/translate", "application/json", byteBody)
	panicIf(err)

	defer res.Body.Close()

	var result interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}

	m := result.(map[string]interface{})

	if val, ok := m["translatedText"]; ok {
		return val.(string)
	}

	if val, ok := m["error"]; ok {
		log.Fatal(val)
	}

	log.Fatal("unknown answer")
	return ""
}

func panicIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
