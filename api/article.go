package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func GetArticle(slug string) (string, error) {
	return slug, nil
}

func Article(w http.ResponseWriter, r *http.Request) {
	// Ensure we have the correct header
	w.Header().Set("Content-Type", "application/json")
	// Get URL param ?nme=file_name
	slug := r.URL.Query().Get("name")
	// Get article from storage
	article, err := GetArticle(strings.ToLower(slug))
	// Handle error
	if err != nil {
		// Set the correct status
		w.WriteHeader(http.StatusNotFound)
		// Log the exception
		log.Printf("Error %s URL %s", err, slug)
		// Notify user
		fmt.Fprintf(w, "Not hotdog")
	}
	// Set the correct status
	w.WriteHeader(http.StatusOK)
	// Return data
	fmt.Fprintf(w, article)
}
