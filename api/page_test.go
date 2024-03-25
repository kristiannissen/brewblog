package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPage200(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/page?name=sample", nil)
	w := httptest.NewRecorder()

	Page(w, req)

	res := w.Result()
	defer res.Body.Close()

	// resp, _ := io.ReadAll(res.Body)
	// log.Println(string(resp))

	if res.StatusCode != 200 {
		t.Errorf("Wrong response %d", res.StatusCode)
	}
}

func TestPage404(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/page?name=", nil)
	w := httptest.NewRecorder()

	Page(w, req)

	res := w.Result()
	defer res.Body.Close()

	// resp, _ := io.ReadAll(res.Body)
	// log.Println(string(resp), res.StatusCode)

	if res.StatusCode != 404 {
		t.Errorf("Wrong response %d", res.StatusCode)
	}
}

func TestPages(t *testing.T) {
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
