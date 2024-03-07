package post

import (
	"fmt"
	"net/http"
)

func Post(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, "Hello Kitty")
}
