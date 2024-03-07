package handler

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestArticle(t *testing.T) {

	req, err := http.NewRequest("GET", "/article", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	h := http.HandlerFunc(Article)
	h.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fatal("Error", status)
	}

	log.Println("Here", rr.Body.String())
}
