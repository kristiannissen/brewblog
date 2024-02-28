package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func GetArticle(slug string) (string, error) {
	var err error
	var p string
	var b []byte

	// Find file
	p, err = os.Getwd()

	if err != nil {
		log.Println(err)
		return "", err
	}
	// Full path
	p = filepath.Join(p + "/../content/" + slug + ".md")
	// Readfile
	b, err = os.ReadFile(p)

	if err != nil {
		log.Println(err)
		return "", err
	}

	// Convert data
	s := string(b)

	return s, nil
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
