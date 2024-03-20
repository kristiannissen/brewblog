package api

import (
	"fmt"
	"net/http"
)

func Pages(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")

	fmt.Fprintf(w, string("Hello"))
}
