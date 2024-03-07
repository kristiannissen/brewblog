package api

import (
	"fmt"
	"log"
	"net/http"
)

func Article(w http.ResponseWriter, r *http.Request) {
	// Ensure we have the correct header
	w.Header().Set("Content-Type", "application/json")
	// Get URL param ?nme=file_name
	slug := r.URL.Query().Get("name")
	// Log the slug
	log.Printf("Query is %s", slug)
	// Get article from storage

	fmt.Fprintf(w, "Hello Kitty")
}
