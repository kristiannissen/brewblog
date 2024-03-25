package pkg

import (
	"strings"
	"testing"

	"brewblog/_pkg/parser"
	"brewblog/_pkg/render"
	s "brewblog/_pkg/service"
	v "brewblog/_pkg/service/vercelservice"
)

func TestPageService(t *testing.T) {
	// Mimics HTTP handlers
	t.Run("404", func(t *testing.T) {
		url, err := PageService("")
		if err == nil {
			t.Errorf("404 not working URL %s", url)
		}
	})

	t.Run("200", func(t *testing.T) {
		url, err := PageService("sample.md")
		if err != nil {
			t.Errorf("200 not working URL %s", url)
		}
	})
}

func TestServiceNew(t *testing.T) {
	provider := s.ServiceProvider(&v.VercelService{})
	t.Errorf("%T", provider)
}

func TestServiceList(t *testing.T) {
	service := s.ServiceProvider(&v.VercelService{})
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
	service := s.ServiceProvider(&v.VercelService{})
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
	service := s.ServiceProvider(&v.VercelService{})
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

## H2 Lorem ippsum
Para 2 Lorem ipsum dolor sit amet, consectetur adipiscing elit.
Sed do eiusmod tempo incididunt ut labore et dolore magna aliqua.

![1 The San Juan!](/1assets/san-juan-mountains.jpg)
![2 The San Juan!](/2assets/san-juan-mountains.jpg)
![3 The San Juan!](/3assets/san-juan-mountains.jpg)

## H2 Lorem ippsum
Para 3 Lorem ipsum dolor sit amet, consectetur adipiscing elit.
Sed do eiusmod tempo incididunt ut labore et dolore magna aliqua.

![4 The San Juan!](/4assets/san-juan-mountains.jpg)
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

		if strings.Index(string(d), "#") != 0 {
			t.Error("Meta not stripped")
		}
	})
}

func TestParseJSON(t *testing.T) {
	m, err := parser.ParseJSON([]byte(d))

	if err != nil {
		t.Error(err)
	}

	if len(m.Title) == 0 {
		t.Error("No title")
	}

	if len(m.Paragraphs) == 0 {
		t.Error("Not all paragraphs found")
	}

}

func TestParseTitle(t *testing.T) {

	t.Run("Title", func(t *testing.T) {
		title := parser.ParseTitle([]byte(d))

		if strings.HasPrefix(title, "H1") == false {
			t.Error("got", title)
		}
	})
}

func TestRenderMarkup(t *testing.T) {
	m := parser.StripMeta([]byte(d))
	m = render.RenderMarkdown(m)

	if strings.HasPrefix(string(m), "<h1>") == false {
		t.Error("Markdown error")
	}
}

func TestRenderTemplate(t *testing.T) {
	m := parser.StripMeta([]byte(d))
	h, _ := render.RenderTemplate(m)

	if strings.HasPrefix(string(h), "<!DOCTYPE") == false {
		t.Error("No template")
	}
}
