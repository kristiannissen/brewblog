package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestArticle(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/articles", nil)
	w := httptest.NewRecorder()
	// Test article handler
	Article(w, req)
	res := w.Result()
	defer res.Body.Close()

	t.Fatal("Not hotdog")
}
