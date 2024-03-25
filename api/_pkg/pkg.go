package pkg

import (
	p "brewblog/_pkg/parser"
	r "brewblog/_pkg/render"
	s "brewblog/_pkg/service"
	v "brewblog/_pkg/service/vercelservice"
)

var service s.BlobService

func init() {
	service = s.ServiceProvider(&v.VercelService{})
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

func Hello() string {
	return "Hello"
}
