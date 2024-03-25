package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	Title string `json:"title"`
}

func Pages(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")

	articles := []Article{
		{Title: "Hello"},
		{Title: "Kitty"},
	}

	b, _ := json.Marshal(articles)

	fmt.Fprintf(w, string(b))
}
