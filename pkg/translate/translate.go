package translate

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"path"

	"log"
	"net/http"
	"net/url"
)

const (
	uri = "http://127.0.0.1:5000"
)

type Config struct {
	Url   string
	Key   string
	Debug io.Writer
}

type Translation struct {
	log *log.Logger
	Config
}

func New(conf Config) *Translation {
	tr := new(Translation)

	if conf.Url != "" {
		tr.Url = conf.Url
	} else {
		tr.Url = uri
	}

	tr.Key = conf.Key

	if conf.Debug == nil {
		conf.Debug = ioutil.Discard
	}

	tr.log = log.New(conf.Debug, "[TranslatePkg]\t", log.LstdFlags)

	return tr
}

func (tr *Translation) Translate(source, sourceLang, targetLang string) (string, error) {
	params := url.Values{}
	params.Set("q", source)
	params.Add("source", sourceLang)
	params.Add("target", targetLang)

	uri, err := url.Parse(tr.Url)
	if err != nil {
		tr.log.Println("Error parse url")
		log.Fatal(err)
	}

	uri.Path = path.Join(uri.Path, "/translate")

	encodedParams := bytes.NewBufferString(params.Encode())

	res, err := http.Post(uri.String(), "application/x-www-form-urlencoded", encodedParams)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Decode the JSON response
	var result interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}

	m := result.(map[string]interface{})

	if val, ok := m["translatedText"]; ok {
		return val.(string), nil
	}

	if val, ok := m["error"]; ok {
		log.Fatal(val)
	}

	return "", log.Output(5, "unknown answer")
}
