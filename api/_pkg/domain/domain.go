package domain

import "time"

type Blob struct {
	URL         string
	PathName    string
	Size        uint64
	UploadedAt  time.Time
	ContentType string
}

type Image struct {
	URL   string `json:"url"`
	Alt   string `json:"alt"`
	Title string `json":title"`
}

type Paragraph struct {
	Body   string  `json:"body,omitempty"`
	Header string  `json:"header,omitempty"`
	Images []Image `json:"images,omitempty"`
}

type Article struct {
	Title      string            `json:"title"`
	Meta       map[string]string `json:"meta"`
	Paragraphs []Paragraph       `json:"paragraphs"`
}
