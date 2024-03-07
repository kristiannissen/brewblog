package post

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Post(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")

	e := Entry{Title: "Kitty"}
	b, _ := json.Marshal(e)

	fmt.Fprintf(w, string(b))
}
