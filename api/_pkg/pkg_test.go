package pkg

import (
	"log"
	"os"
	"strings"
	"testing"

	"brewblog/_pkg/parser"
	"brewblog/_pkg/render"
	s "brewblog/_pkg/service"
	v "brewblog/_pkg/service/vercelservice"
)

func TestPageService(t *testing.T) {
	// Get list
	l, _ := service.List()

	url, err := PageService(l[0].PathName)
	if err != nil {
		t.Errorf("200 not working URL %s", url)
	}
}

func TestPageRecentService(t *testing.T) {
	_, err := PageRecentService()

	if err != nil {
		t.Error(err)
	}
}

func TestPagesService(t *testing.T) {
	_, err := PagesService()
	if err != nil {
		t.Errorf("Pages %s", err)
	}
}

func TestServiceNew(t *testing.T) {
	t.Errorf("%T", s.ServiceProvider(&v.VercelService{}))
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
	l, _ := service.List()

	bytes, err := service.Download(l[0].URL)

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
URL: sample.md
title: Lorem ipsum dolor sit amet
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

func TestParseStruct(t *testing.T) {
	var b []byte
	var e error
	if b, e = os.ReadFile("../../content/batch_231.md"); e != nil {
		log.Fatal(e)
	}

	// TODO: Rename to renderStruct
	m, err := parser.ParseJSON([]byte(b))
	// log.Println(m)

	if err != nil {
		t.Error(err)
	}

	if len(m.Title) == 0 {
		t.Error("No title")
	}

	if len(m.Paragraphs) == 0 {
		t.Error("No paragraphs found")
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
