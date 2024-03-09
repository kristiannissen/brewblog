package post

import (
	"encoding/json"
	"fmt"
	"net/http"

	p "brewblog/_pkg"
)

type Entry struct {
	Title string `json:"title"`
}

func Post(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")

	e := Entry{Title: p.Hello()}
	b, _ := json.Marshal(e)

	fmt.Fprintf(w, string(b))
}
