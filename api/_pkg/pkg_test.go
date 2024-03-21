package pkg

import (
	"strings"
	"testing"

	parser "brewblog/_pkg/parser"
	s "brewblog/_pkg/service"
)

func TestServiceNew(t *testing.T) {
	service := s.NewVercelService()

	t.Error("Yoko Oh No", service)
}

func TestServiceList(t *testing.T) {
	service := s.NewVercelService()
	list, err := service.List()

	t.Run("No error", func(t *testing.T) {
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("List len", func(t *testing.T) {
		if len(list) == 0 {
			t.Error("Len is 0")
		}
	})
}

func TestServiceFind(t *testing.T) {
	service := s.NewVercelService()
	url, err := service.Find("sample.md")

	t.Run("Test Error", func(t *testing.T) {
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Test URL", func(t *testing.T) {
		if url == "" {
			t.Error("No URL")
		}
	})
}

func TestServiceDownload(t *testing.T) {
	service := s.NewVercelService()
	url, _ := service.Find("sample.md")

	if url == "" {
		t.Error("No URL Found")
	}

	bytes, err := service.Download(url)

	if err != nil {
		t.Error(err)
	}

	if len(bytes) == 0 {
		t.Error("Nothing to download")
	}
}

var d = `---
tags: Hello Kitty
published: 2024-02-29
URL: /hello-kitty
---

# H1 Lorem ipsum dolor sit amet

Para 1 Lorem ipsum dolor sit amet, consectetur adipiscing elit.
Sed do eiusmod tempo incididunt ut labore et dolore magna aliqua.
`

func TestMeta(t *testing.T) {
	h := parser.ParseMeta([]byte(d))

	t.Run("Has tags", func(t *testing.T) {
		if h["tags"] == "" {
			t.Error("No tags found")
		}
	})

	t.Run("Strip meta", func(t *testing.T) {
		d := parser.StripMeta([]byte(d))

		if strings.HasPrefix(string(d), "# H1") == false {
			t.Error("Meta not stripped")
		}
	})
}

func TestParseTitle(t *testing.T) {

	t.Run("Title", func(t *testing.T) {
		title := parser.ParseTitle([]byte(d))

		if strings.HasPrefix(title, "H1") == false {
			t.Error("got", title)
		}
	})
}
