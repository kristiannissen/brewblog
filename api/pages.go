package api

import (
	p "brewblog/_pkg"
	"net/http"
)

func Pages(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")

	l, err := p.PagesService()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error"))
	} else {
		var resp string
		resp = `{"pages": ` + string(l) + `}`
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp))
	}
}
