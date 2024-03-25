package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRecent(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/recent", nil)
	w := httptest.NewRecorder()

	Recent(w, req)

	res := w.Result()
	defer res.Body.Close()

	// resp, _ := io.ReadAll(res.Body)
	// log.Println(string(resp))

	if res.StatusCode != 200 {
		t.Errorf("Wrong response %d", res.StatusCode)
	}

}

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
	// log.Println(string(data))

	if strings.HasPrefix(string(data), "{") == false {
		t.Fatal("Empty list")
	}
}
