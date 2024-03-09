package post

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Posts(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")

	var l []Entry

	e := Entry{Title: "Kitty"}

	l = append(l, e)

	b, _ := json.Marshal(l)

	fmt.Fprintf(w, string(b))
}
