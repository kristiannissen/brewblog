package user

import (
	"fmt"
	"net/http"
)

type UserObj struct {
	Name string
}

func User(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")

	fmt.Fprintf(w, "Hello")
}
