package post

import (
	"log"
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

	log.Println(res)

	t.Fatal("oh no")
}
