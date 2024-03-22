package domain

import "time"

// TODO: Move to seperate folder
type Blob struct {
	URL         string
	PathName    string
	Size        uint64
	UploadedAt  time.Time
	ContentType string
}

type Image struct {
	URL   string
	Alt   string
	Title string
}

type Paragraph struct {
	Body   string
	Header string
	Images []Image
}

type Article struct {
	Title      string            `json:"title"`
	Meta       map[string]string `json:"meta"`
	Paragraphs []Paragraph
}
