package handler

import (
	"fmt"
	"net/http"
)

func Articles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "articles")
}
