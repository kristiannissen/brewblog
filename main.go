package main

import (
	"brewblog/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/articles", api.Articles)

	log.Fatal(http.ListenAndServe(":9009", nil))
}
