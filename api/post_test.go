package post

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPost(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/post", nil)
	w := httptest.NewRecorder()

	Post(w, req)

	res := w.Result()
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	e := Entry{}
	_ = json.Unmarshal(data, &e)

	if e.Title != "Hello" {
		t.Fatalf("Error: want Kitty got %s", e.Title)
	}
}

func TestPosts(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/posts", nil)
	w := httptest.NewRecorder()

	Posts(w, req)

	res := w.Result()
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)

	if string(data) != "Hello" {
		t.Fatal("Empty list")
	}
}
