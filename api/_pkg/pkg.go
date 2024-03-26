package pkg

import (
	domain "brewblog/_pkg/domain"
	p "brewblog/_pkg/parser"
	r "brewblog/_pkg/render"
	s "brewblog/_pkg/service"
	v "brewblog/_pkg/service/vercelservice"
	"encoding/json"
)

var service s.BlobService

func init() {
	service = s.ServiceProvider(&v.VercelService{})
}

func PageRecentService() ([]byte, error) {
	var b []byte
	var err error
	var l []domain.Blob

	l, err = service.List()
	if err != nil {
		return b, err
	}
	//
	b, err = service.Download(l[0].URL)
	if err != nil {
		return b, err
	}
	m := p.StripMeta(b)
	var a domain.Article
	a, err = p.ParseJSON(m)
	b, err = json.Marshal(a)

	return b, err
}

func PageService(name string) ([]byte, error) {
	// Find fully qualified URL
	var err error
	var b []byte
	var url string

	url, err = service.Find(name)

	if err != nil {
		// URL not found
		return b, err
	}
	// Download data
	b, err = service.Download(url)

	if err != nil {
		// Error in download
		return b, err
	}
	// Build markup
	m := p.StripMeta(b)
	m = r.RenderMarkdown(m)
	h, _ := r.RenderTemplate(m)

	return []byte(h), nil
}

type Page struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Image string `json:"image"`
}

func PagesService() ([]byte, error) {
	var err error
	var l []domain.Blob
	var b []byte
	// List of pages
	l, err = service.List()

	if err != nil {
		return b, err
	}

	var pages []Page

	for _, v := range l {
		var doc []byte
		doc, err = service.Download(v.URL)
		if err != nil {
			break
		}
		m := p.ParseMeta(doc)
		pages = append(pages,
			Page{Title: m["title"], URL: "/api/page?name=" + v.PathName, Image: "https://placehold.co/600x400"},
		)
	}

	b, err = json.Marshal(pages)

	return b, err
}
