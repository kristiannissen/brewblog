package handler

import (
	"fmt"
	"net/http"
)

func GetArticle(slug string) string {
	return slug
}

func Article(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "article")
}
