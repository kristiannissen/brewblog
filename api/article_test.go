package handler

import (
	"testing"
)

func TestGetArticle(t *testing.T) {
	a, _ := GetArticle("hello")

	if a != "hello" {
		t.Fatal("not hotdog")
	}
}
