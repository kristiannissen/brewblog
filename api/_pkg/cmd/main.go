package main

import (
	api "brewblog"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/api/recent", api.Recent)
	http.HandleFunc("/api/pages", api.Pages)
	http.HandleFunc("/api/page", api.Page)

	log.Fatal(http.ListenAndServe(":9001", nil))
}
