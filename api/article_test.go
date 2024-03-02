package handler

import (
	"testing"
)

func TestGetArticle(t *testing.T) {
	_, err := GetArticle("sample")

	if err != nil {
		t.Fatal(err)
	}
}
