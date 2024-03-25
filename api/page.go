package api

import (
	"fmt"
	"net/http"

	p "brewblog/_pkg"
)

func Page(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	n := req.URL.Query().Get("name")
	// Get page data
	resp, err := p.PageService(n)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - not found"))
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(resp))
	}
}
