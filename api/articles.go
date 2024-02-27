package handler

import (
	"encoding/json"
	"net/http"
)

var ArticleList = []Article{
	{Title: "Hello Kitty", Slug: "hello-kitty", Paragraphs: []Paragraph{
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
