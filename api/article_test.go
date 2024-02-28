package handler

import (
	"log"
	"testing"
)

func TestGetArticle(t *testing.T) {
	a, _ := GetArticle("sample")

	log.Println(a)

	if a != "hello" {
		t.Fatal("not hotdog")
	}
}
