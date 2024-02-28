package handler

import (
	"testing"
)

func TestGetArticle(t *testing.T) {
	a := GetArticle("hello")

	if a != "hello" {
		t.Fatal("not hotdog")
	}
}
