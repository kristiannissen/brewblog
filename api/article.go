package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Entry struct {
	Title string `json:"title"`
	Intro string `json:"intro"`
	Slug  string `json:"slug"`
}

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
	p = filepath.Join(p + "/_content/" + slug + ".md")
	// Readfile
	b, err = os.ReadFile(p)

	if err != nil {
		log.Println(err)
		return "", err
	}

	// Convert data
	e, _ := json.Marshal(Transform(string(b)))

	return string(e), nil
}

func Transform(s string) Entry {
	// Create a new Entry
	e := Entry{}

	return e
}

func Article(w http.ResponseWriter, r *http.Request) {
	// Ensure we have the correct header
	w.Header().Set("Content-Type", "application/json")
	// Get URL param ?nme=file_name
	slug := r.URL.Query().Get("name")
	// Log the slug
	log.Printf("Query is %s", slug)
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
	} else {
		// Set the correct status
		w.WriteHeader(http.StatusOK)
		// Return data
		fmt.Fprintf(w, article)
	}
}
