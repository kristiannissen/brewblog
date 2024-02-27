package handler

import (
	"testing"
)

func TestArticle(t *testing.T) {
	_, err := GetArticle()

	// log.Println(article)

	if err != nil {
		t.Fatal(err)
	}
}
