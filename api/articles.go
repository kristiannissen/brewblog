package handler

import (
	"encoding/json"
	"net/http"
)

type Paragraph struct {
	Header    string `json:"header"`
	Paragraph string `json:"paragraph"`
}

type Article struct {
	Title      string      `json:"title"`
	Intro      string      `json:"intro"`
	Paragraphs []Paragraph `json:"paragraphs"`
	Slug       string      `json:"slug"`
}

var ArticleList = []Article{
	{Title: "Hello Kitty", Intro: "Hello kitty", Slug: "hello-kitty", Paragraphs: []Paragraph{
		{Header: "Hello Kitty", Paragraph: "Hello Kitty"},
		{Header: "Hello Kitty", Paragraph: "Hello Kitty"},
	}},
}

func Articles(w http.ResponseWriter, r *http.Request) {
	// Add proper headers
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json, _ := json.Marshal(ArticleList)

	w.Write(json)
}

func GetArticles() (string, error) {
	// Return json
	json, err := json.Marshal(ArticleList)

	return string(json), err
}
