package api

import (
	"fmt"
	"net/http"
)

func Page(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	n := req.URL.Query().Get("name")

	fmt.Fprintf(w, n)
}
