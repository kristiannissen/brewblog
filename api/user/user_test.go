package user

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/post", nil)
	w := httptest.NewRecorder()

	User(w, req)

	res := w.Result()
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)

	t.Fatalf("Oh no! %s", data)
}
