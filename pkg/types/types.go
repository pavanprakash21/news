package types

import (
	"time"
)

type Article struct {
	Source struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Urltoimage  string    `json:"urlToImage"`
	Publishedat time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}

type NewsResult struct {
	Status   string    `json:"status"`
	Topic    string    `json:"topic"`
	Articles []Article `json:"articles"`
}

type FinalData struct {
	Data Data `json:"data"`
}

type Data struct {
	News []NewsResult `json:"news"`
}
