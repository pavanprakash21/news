package concurrently

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"

	"github.com/pavanprakash21/news/pkg/types"
)

func MakeRequests(urls []string, concurrencyLimit uint8) []types.NewsResult {
	type respMap struct {
		index    int
		response http.Response
		err      error
	}

	semaphoreChan := make(chan struct{}, concurrencyLimit)
	resultsChan := make(chan *respMap)

	defer func() {
		close(semaphoreChan)
		close(resultsChan)
	}()

	for idx, url := range urls {
		go func(idx int, url string) {
			semaphoreChan <- struct{}{}
			res, err := http.Get(url)
			panicIf(err)
			resp := &respMap{idx, *res, err}
			resultsChan <- resp
			<-semaphoreChan
		}(idx, url)
	}

	var resps []respMap

	for {
		resp := <-resultsChan
		resps = append(resps, *resp)

		if len(resps) == len(urls) {
			break
		}
	}

	sort.Slice(resps, func(i, j int) bool {
		return resps[i].index < resps[j].index
	})

	var results []types.NewsResult

	for _, resp := range resps {
		body, err := ioutil.ReadAll(resp.response.Body)
		panicIf(err)
		var res types.NewsResult
		data := []byte(body)
		e := json.Unmarshal(data, &res)
		panicIf(e)
		results = append(results, res)
	}

	return results
}

func panicIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
