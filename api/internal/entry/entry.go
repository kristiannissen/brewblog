// Package entry exposes the entry struct
// and offers transformation features
// that turns markdown into a json structure

package entry

import (
	"encoding/json"
	"regexp"
	"strings"
)

const (
	SEP    = "---"
	IMG    = "!"
	HEADER = "#"
	LB     = "\n"
	DLB    = "\n\n"
)

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Paragraph struct {
	Header string  `json:"header"`
	Body   string  `json:"body"`
	Images []Image `json:"images"`
}

type Image struct {
	URL  string `json:"url"`
	Text string `json:"text"`
}

type Entry struct {
	Title      string      `json:"title"`
	Meta       []KeyValue  `json:"meta"`
	Paragraphs []Paragraph `json:"paragraphs"`
}

func GetEntry(n string) (Entry, error) {
	e := Entry{Title: "Hello Kitty"}

	return e, nil
}

func extractMeta(s string) []KeyValue {
	var data []KeyValue

	// Split string
	p := strings.Split(s, "\n")

	for _, v := range p {
		// Split into key/value
		d := strings.Split(v, ":")
		if d[0] != "" {
			kv := KeyValue{}
			kv.Key = d[0]
			kv.Value = d[1]
			data = append(data, kv)
		}
	}

	return data
}

func extractHeader(s string) string {
	return strings.TrimSpace(
		s[strings.LastIndex(s, HEADER)+1:])
}

func extractImage(s string) Image {
	img := Image{}

	img.URL = s[strings.Index(s, "(")+1 : strings.LastIndex(s, ")")]
	img.Text = s[strings.Index(s, "[")+1 : strings.LastIndex(s, "]")]

	return img
}

func ParseEntryData(s string) (Entry, error) {
	e := Entry{}
	var f, l int
	f = 0
	l = 0
	// Extract meta data
	if strings.HasPrefix(s, SEP) {
		// Read string until ending ---
		f = strings.Index(s, SEP) + len(SEP)
		l = strings.Index(s[f:], SEP) + len(SEP)
		// Append to Entry
		e.Meta = extractMeta(s[f:l])
		// Adjust string
		f = l + len(SEP)
		s = s[f:]
	}

	// Extract title (h1)
	re := regexp.MustCompile(`#\s(.*)?`)
	loc := re.FindIndex([]byte(s))

	if len(loc) > 0 {
		e.Title = extractHeader(s[loc[0]:loc[1]])
		// Adjust f
		f = loc[1]
		// Adjust string
		s = s[f:]
	}

	// Split into parts
	var parts []string
	parts = strings.Split(s, DLB)
	for _, part := range parts {
		// Create a new paragraph
		para := Paragraph{}
		// Extract title (H2, h3)
		if strings.HasPrefix(part, HEADER) {
			// Headline
			para.Header = extractHeader(part)
		} else if strings.HasPrefix(part, IMG) {
			// Images
			for _, img := range strings.Split(part, LB) {
				para.Images = append(para.Images, extractImage(img))
			}
		} else {
			// Good old text
			para.Body = part
		}
		// Add to array
		if part != "" {
			e.Paragraphs = append(e.Paragraphs, para)
		}
	}

	return e, nil
}

func EntryToJson(e Entry) (string, error) {
	s, err := json.Marshal(e)

	return string(s), err
}
