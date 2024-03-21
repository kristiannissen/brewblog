package pkg

import (
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

	if err != nil {
		t.Error(err)
	}

	if len(list) == 0 {
		t.Error("No blobs")
	}
}

func TestServiceFind(t *testing.T) {
	service := s.NewVercelService()
	url, err := service.Find("sample.md")

	if err != nil {
		t.Error(err)
	}

	if url == "" {
		t.Error("No URL found")
	}
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

func TestParseMeta(t *testing.T) {
	h := parser.ParseMeta([]byte(d))

	if h["tags"] != "Hello Kitty" {
		t.Error("Parser is not working")
	}
}
