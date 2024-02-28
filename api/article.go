package handler

import (
	"fmt"
	"net/http"
)
/*
type Paragraph struct {
	Header    string `json:"header"`
	Paragraph string `json:"paragraph"`
}

type Article struct {
	Title      string      `json:"title"`
	Paragraphs []Paragraph `json:"paragraphs"`
	Slug       string      `json:"slug"`
}

// TODO: should take file name as parameter
func GetArticle() (Article, error) {
	data, err := os.ReadFile("../content/sample.md")
	if err != nil {
		panic(err)
	}

	// Split text into chunks
	chunks := strings.Split(string(data), "\n\n")
	article := Article{}
	// Populate article
	article.Title = strings.TrimPrefix(chunks[0], "# ")
	// Pass rest into paragraphs
	for _, v := range chunks[1:] {
		para := Paragraph{}
		if strings.HasPrefix(v, "#") {
			para.Header = v
		} else {
			para.Paragraph = v
		}
		article.Paragraphs = append(article.Paragraphs, para)
	}

	return article, err
}
*/
// Handler needed
func Article(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello Kitty")
}
