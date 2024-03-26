package api

import (
	"fmt"
	"net/http"

	p "brewblog/_pkg"
)

func Recent(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Cache-Control", "max-age=0, s-maxage=86400")
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Get page data
	resp, err := p.PageRecentService()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - not found"))
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(resp))
	}
}
