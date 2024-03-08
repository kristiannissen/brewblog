package post

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	u "brewblog/user"
)

type Entry struct {
	Title string `json:"title"`
}

func Post(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")

	e := Entry{Title: "Kitty"}
	b, _ := json.Marshal(e)

	user := u.UserObj{}
	log.Println(user)

	fmt.Fprintf(w, string(b))
}

func Posts(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")

	var l []Entry

	e := Entry{Title: "Kitty"}

	l = append(l, e)

	b, _ := json.Marshal(l)

	fmt.Fprintf(w, string(b))
}
