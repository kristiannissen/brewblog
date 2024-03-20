package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/page", nil)
	w := httptest.NewRecorder()

	Page(w, req)

	res := w.Result()
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)

	if string(data) != "Hello" {
		t.Fatal("Not Hello")
	}
}

func TestPagess(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/pages", nil)
	w := httptest.NewRecorder()

	Pages(w, req)

	res := w.Result()
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)

	if string(data) != "Hello" {
		t.Fatal("Empty list")
	}
}
