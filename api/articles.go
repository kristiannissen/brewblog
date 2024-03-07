package handler

import (
	"_internal/entry"
	"fmt"
	"log"
	"net/http"
)

func Articles(w http.ResponseWriter, r *http.Request) {
	l := entry.GetEntries()
	log.Println(l)

	fmt.Fprintf(w, "articles")
}
