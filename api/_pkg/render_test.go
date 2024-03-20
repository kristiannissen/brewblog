package pkg

import (
	"log"
	"testing"
)

var d = `---
tags: Hello Kitty
published: 2024-02-29
URL: /hello-kitty
---

# H1 Lorem ipsum dolor sit amet

Para 1 Lorem ipsum dolor sit amet, consectetur adipiscing elit.
Sed do eiusmod tempo incididunt ut labore et dolore magna aliqua.
`

func TestParseMeta(t *testing.T) {
	h, err := ParseMeta([]byte(d))

	if err != nil {
		t.Fatal(err)
	}

	if len(h) == 0 {
		t.Fatal("Map is empty")
	}

	if h["tags"] != "Hello Kitty" {
		t.Fatal("Tags wrong")
	}
}

func TestRenderMarkdown(t *testing.T) {
	r, _ := RenderMarkdown([]byte(d))
	log.Println(string(r))

	t.Fatal("Oh no Yoko")
}

func TestRenderHTML(t *testing.T) {
	md, _ := RenderMarkdown([]byte(d))
	html, _ := RenderHTML(md)

	if len(html) == 0 {
		t.Fatal("HTML is empty")
	}
}
