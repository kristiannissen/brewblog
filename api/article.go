package handler

import (
  "fmt"
  "net/http"
)

func Article(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "vim-go")
}
