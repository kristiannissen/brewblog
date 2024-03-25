package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Image string `json:"image"`
}

func Pages(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")

	articles := []Article{
		{Title: "Hello", URL: "/post/", Image: "https://placehold.co/600x400"},
		{Title: "Kitty", URL: "/post", Image: "https://placehold.co/600x400"},
	}

	b, _ := json.Marshal(articles)
	var resp string
	resp = `{"pages": ` + string(b) + `}`

	fmt.Fprintf(w, resp)
}
