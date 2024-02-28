package handler

import (
	"fmt"
	"net/http"
)

func GetArticle(string slug) string {
	return "Hello Kitty"
}

func Article(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "article")
}
