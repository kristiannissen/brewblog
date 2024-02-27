package handler

import (
	"testing"
)

func TestArticles(t *testing.T) {
	// GetArticles
	_, err := GetArticles()

	// log.Println(articles)

	if err != nil {
		t.Fatal(err)
	}
}
