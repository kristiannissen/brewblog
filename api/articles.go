package api

import (
	"brewblog/api/internal/entry"
	"fmt"
	"log"
	"net/http"
)

func Articles(w http.ResponseWriter, r *http.Request) {
	l, err := entry.GetEntries()
	log.Println(l, err)

	fmt.Fprintf(w, "articles")
}
